package conekta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"errors"
)

// Conekta API main endpoint
const baseUrl = "https://api.conekta.io/"

// Available configuration options, if not provided sane values will be
// used by default
type Options struct {
	// Time to wait for requests, in seconds
	Timeout uint

	// Time to maintain open the connection with the service, in seconds
	KeepAlive uint

	// Maximum network connections to keep open with the service
	MaxConnections uint

	// API version to use
	APIVersion string

	// User agent value to report to the service
	UserAgent string
}

// Main service handler
type Client struct {
	// Methods related to 'orders' management
	Orders *ordersClient

	// Methods related to 'customers' management
	Customers *customersClient

	// Methods related to 'plans' management
	Plans *plansClient

	c          *http.Client
	key        string
	apiVersion string
	userAgent  string
}

// Network request options
type requestOptions struct {
	method   string
	endpoint string
	data     interface{}
}

// Return sane default configuration values
func defaultOptions() *Options {
	return &Options{
		Timeout:        30,
		KeepAlive:      600,
		MaxConnections: 100,
		APIVersion:     "v2.0.0",
		UserAgent:      "",
	}
}

// New will construct a usable service handler using the provided API key and
// configuration options, if 'nil' options are provided default sane values will
// be used
func New(key string, options *Options) (*Client, error) {
	if key == "" {
		return nil, errors.New("API key is required")
	}
	
	// If no options are provided, use default sane values
	if options == nil {
		options = defaultOptions()
	}

	// Configure base HTTP transport
	t := &http.Transport{
		MaxIdleConns:        int(options.MaxConnections),
		MaxIdleConnsPerHost: int(options.MaxConnections),
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(options.Timeout) * time.Second,
			KeepAlive: time.Duration(options.KeepAlive) * time.Second,
			DualStack: true,
		}).DialContext,
	}

	// Setup main client
	client := &Client{
		key:        key,
		apiVersion: options.APIVersion,
		userAgent:  options.UserAgent,
		c: &http.Client{
			Transport: t,
			Timeout:   time.Duration(options.Timeout) * time.Second,
		},
	}
	client.Orders = &ordersClient{c: client}
	client.Customers = &customersClient{c: client}
	client.Plans = &plansClient{c: client}
	return client, nil
}

// Dispatch a network request to the service
func (i *Client) request(r *requestOptions) ([]byte, error) {
	// Build request with headers and credentials
	data, _ := json.Marshal(r.data)
	req, _ := http.NewRequest(r.method, r.endpoint, bytes.NewReader(data))
	req.Header.Add("Accept", fmt.Sprintf("application/vnd.conekta-%s+json", i.apiVersion))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(i.key, "")
	if i.userAgent != "" {
		req.Header.Add("User-Agent", i.userAgent)
	}

	// Execute request
	res, err := i.c.Do(req)
	if res != nil {
		// Properly discard request content to be able to reuse the connection
		defer io.Copy(ioutil.Discard, res.Body)
		defer res.Body.Close()
	}

	// Network level errors
	if err != nil {
		return nil, err
	}

	// Get response contents
	body, err := ioutil.ReadAll(res.Body)

	// Application level errors
	if res.StatusCode != 200 {
		e := &APIError{}
		json.Unmarshal(body, e)
		return nil, e
	}
	return body, nil
}
