package conekta

import (
	"encoding/json"
	"net/http"
	"path"
)

// Defines the public interface required to access available 'customers' methods
type CustomersAPI interface {
	// Creates a new customer
	// https://developers.conekta.com/api?language=bash#create-customer
	Create(customer *Customer) error

	// Updates a existing customer
	// https://developers.conekta.com/api?language=bash#update-customer
	Update(customer *Customer) error

	// Deletes a existing customer
	// https://developers.conekta.com/api?language=bash#capture-order
	Delete(customerID string) error

	// Creates new payment source
	// https://developers.conekta.com/api?language=bash#payment-source
	CreatePaymentSource(customerID, tokenID string) error

	// Updates existing payment source
	// https://developers.conekta.com/api?language=bash#update-payment-source
	UpdatePaymentSource(customerID string, update *PaymentSourceUpdate) error

	// Deletes existing payment source
	// https://developers.conekta.com/api?language=bash#delete-payment-source
	DeletePaymentSource(customerID, sourceID string) error

	// Creates a new Shipping Contact for an existing customer
	// https://developers.conekta.com/api?language=bash#create-shipping-contact-customer
	CreateShippingContact(customerID string, contact *ShippingContact) error

	// Updates an existing Shipping Contact
	// https://developers.conekta.com/api?language=bash#update-shipping-contact
	UpdateShippingContact(customerID string, contact *ShippingContact) error

	// Deletes an existing Shipping Contact
	// https://developers.conekta.com/api?language=bash#update-shipping-contact
	DeleteShippingContact(customerID, contactID string) error

	// Creates a new subscription using tokenized data
	// https://developers.conekta.com/api?language=bash#create-subscription
	CreateSubscription(customer *Customer, planID, cardID string) error

	// Updates a subscription with a different card or plan
	// https://developers.conekta.com/api?language=bash#update-subscription
	UpdateSubscription(customer *Customer, planID, cardID string) error

	// Pauses a subscription
	// https://developers.conekta.com/api?language=bash#pause-subscription
	PauseSubscription(customerID, subscriptionID string) error

	// Resume a subscription
	// https://developers.conekta.com/api?language=bash#resume-subscription
	ResumeSubscription(customerID, subscriptionID string) error

	// Cancel a subscription
	// https://developers.conekta.com/api?language=bash#resume-subscription
	CancelSubscription(customerID, subscriptionID string) error
}

type customersClient struct {
	c *Client
}

func (cc *customersClient) Create(customer *Customer) error {
	b, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + "customers",
		method:   http.MethodPost,
		data:     customer,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(b, customer)
	return nil
}

func (cc *customersClient) Update(customer *Customer) error {
	b, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customer.ID),
		method:   http.MethodPut,
		data:     customer,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(b, customer)
	return nil
}

func (cc *customersClient) Delete(customerID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID),
		method:   http.MethodDelete,
		data:     customerID,
	})
	return err
}

func (cc *customersClient) CreatePaymentSource(customerID, tokenID string) error {
	data := map[string]string{
		"type":     "card",
		"token_id": tokenID,
	}
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "payment_sources"),
		method:   http.MethodPost,
		data:     data,
	})
	if err != nil {
		return err
	}
	return nil
}

func (cc *customersClient) UpdatePaymentSource(customerID string, update *PaymentSourceUpdate) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "payment_sources", update.ID),
		method:   http.MethodPut,
		data:     update,
	})
	return err
}

func (cc *customersClient) DeletePaymentSource(customerID, sourceID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "payment_sources", sourceID),
		method:   http.MethodDelete,
		data:     customerID,
	})
	return err
}

func (cc *customersClient) CreateShippingContact(customerID string, contact *ShippingContact) error {
	b, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "shipping_contacts"),
		method:   http.MethodPost,
		data:     contact,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(b, contact)
	return nil
}

func (cc *customersClient) UpdateShippingContact(customerID string, contact *ShippingContact) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "shipping_contacts", contact.ID),
		method:   http.MethodPut,
		data:     contact,
	})
	if err != nil {
		return err
	}
	return nil
}

func (cc *customersClient) DeleteShippingContact(customerID, contactID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "shipping_contacts", contactID),
		method:   http.MethodDelete,
		data:     customerID,
	})
	return err
}

func (cc *customersClient) CreateSubscription(customer *Customer, planID, cardID string) error {
	data := map[string]string{"plan": planID}
	if cardID != "" {
		data["card"] = cardID
	}
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customer.ID, "subscription"),
		method:   http.MethodPost,
		data:     data,
	})
	return err
}

func (cc *customersClient) UpdateSubscription(customer *Customer, planID, cardID string) error {
	data := map[string]string{"plan": planID}
	if cardID != "" {
		data["card"] = cardID
	}
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customer.ID, "subscription"),
		method:   http.MethodPut,
		data:     data,
	})
	return err
}

func (cc *customersClient) PauseSubscription(customerID, subscriptionID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "subscription", "pause"),
		method:   http.MethodPost,
		data:     map[string]string{"id": subscriptionID},
	})
	return err
}

func (cc *customersClient) ResumeSubscription(customerID, subscriptionID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "subscription", "resume"),
		method:   http.MethodPost,
		data:     map[string]string{"id": subscriptionID},
	})
	return err
}

func (cc *customersClient) CancelSubscription(customerID, subscriptionID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "subscription", "cancel"),
		method:   http.MethodPost,
		data:     map[string]string{"id": subscriptionID},
	})
	return err
}
