package conekta

import (
	"encoding/json"
	"net/http"
	"path"
)

type customersClient struct {
	c *Client
}

// Creates a new customer
// https://developers.conekta.com/api?language=bash#create-customer
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

// Updates a existing customer
// https://developers.conekta.com/api?language=bash#update-customer
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

// Deletes a existing customer
// https://developers.conekta.com/api?language=bash#capture-order
func (cc *customersClient) Delete(customerID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID),
		method:   http.MethodDelete,
		data:     customerID,
	})
	return err
}

// Creates new payment source
// https://developers.conekta.com/api?language=bash#payment-source
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

// Updates existing payment source
// https://developers.conekta.com/api?language=bash#update-payment-source
func (cc *customersClient) UpdatePaymentSource(customerID string, update *PaymentSourceUpdate) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "payment_sources", update.ID),
		method:   http.MethodPut,
		data:     update,
	})
	return err
}

// Deletes existing payment source
// https://developers.conekta.com/api?language=bash#delete-payment-source
func (cc *customersClient) DeletePaymentSource(customerID, sourceID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "payment_sources", sourceID),
		method:   http.MethodDelete,
		data:     customerID,
	})
	return err
}

// Creates a new Shipping Contact for an existing customer
// https://developers.conekta.com/api?language=bash#create-shipping-contact-customer
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

// Updates an existing Shipping Contact
// https://developers.conekta.com/api?language=bash#update-shipping-contact
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

// Deletes an existing Shipping Contact
// https://developers.conekta.com/api?language=bash#update-shipping-contact
func (cc *customersClient) DeleteShippingContact(customerID, contactID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "shipping_contacts", contactID),
		method:   http.MethodDelete,
		data:     customerID,
	})
	return err
}

// Creates a new subscription using tokenized data
// https://developers.conekta.com/api?language=bash#create-subscription
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

// Updates a subscription with a different card or plan
// https://developers.conekta.com/api?language=bash#update-subscription
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

// Pauses a subscription
// https://developers.conekta.com/api?language=bash#pause-subscription
func (cc *customersClient) PauseSubscription(customerID, subscriptionID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "subscription", "pause"),
		method:   http.MethodPost,
		data:     map[string]string{"id": subscriptionID},
	})
	return err
}

// Resume a subscription
// https://developers.conekta.com/api?language=bash#resume-subscription
func (cc *customersClient) ResumeSubscription(customerID, subscriptionID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "subscription", "resume"),
		method:   http.MethodPost,
		data:     map[string]string{"id": subscriptionID},
	})
	return err
}

// Cancel a subscription
// https://developers.conekta.com/api?language=bash#resume-subscription
func (cc *customersClient) CancelSubscription(customerID, subscriptionID string) error {
	_, err := cc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("customers", customerID, "subscription", "cancel"),
		method:   http.MethodPost,
		data:     map[string]string{"id": subscriptionID},
	})
	return err
}
