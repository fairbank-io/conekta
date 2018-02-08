package conekta

import (
	"path"
	"net/http"
	"encoding/json"
)

type ordersClient struct {
	c *Client
}

// Creates a new order
// https://developers.conekta.com/api?language=bash#create-order
func (oc *ordersClient) Create(order *Order) error {
	b, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + "orders",
		method:   http.MethodPost,
		data:     order,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(b, order)
	return nil
}

// Updates a existing order
// https://developers.conekta.com/api?language=bash#update-order
func (oc *ordersClient) Update(order *Order) error {
	b, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", order.ID),
		method:   http.MethodPut,
		data:     order,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(b, order)
	return nil
}

// Process a pre-authorized order
// https://developers.conekta.com/api?language=bash#capture-order
func (oc *ordersClient) Capture(orderID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "capture"),
		method:   http.MethodPost,
		data:     map[string]string{orderID: orderID},
	})
	if err != nil {
		return err
	}
	return nil
}

// A Refund details the amount and reason why an order was refunded
// https://developers.conekta.com/api?language=bash#refund-order
func (oc *ordersClient) Refund(orderID string, r *Refund) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "refunds"),
		method:   http.MethodPost,
		data:     r,
	})
	if err != nil {
		return err
	}
	return nil
}

// Create a new line item
// https://developers.conekta.com/api?language=bash#create-line-item
func (oc *ordersClient) CreateLineItem(orderID string, item *LineItem) (string, error) {
	res, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "line_items"),
		method:   http.MethodPost,
		data:     item,
	})
	if err != nil {
		return "", err
	}

	var info = map[string]interface{}{}
	json.Unmarshal(res, &info)
	return info["id"].(string), nil
}

// Updates a line item
// https://developers.conekta.com/api?language=bash#update-line-item
func (oc *ordersClient) UpdateLineItem(orderID string, item *LineItem) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "line_items", item.ID),
		method:   http.MethodPut,
		data:     item,
	})
	if err != nil {
		return err
	}
	return nil
}

// Deletes a line item
// https://developers.conekta.com/api?language=bash#delete-line-item
func (oc *ordersClient) DeleteLineItem(orderID, itemID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "line_items", itemID),
		method:   http.MethodDelete,
		data:     itemID,
	})
	return err
}