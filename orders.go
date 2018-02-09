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

// Creates a new Discount Line
// https://developers.conekta.com/api?language=bash#create-discount-line
func (oc *ordersClient) CreateDiscountLine(orderID string, discount *DiscountLine) (string, error) {
	res, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "discount_lines"),
		method:   http.MethodPost,
		data:     discount,
	})
	if err != nil {
		return "", err
	}

	var info = map[string]interface{}{}
	json.Unmarshal(res, &info)
	return info["id"].(string), nil
}

// Updates an existing Discount Line
// https://developers.conekta.com/api?language=bash#update-discount-line
func (oc *ordersClient) UpdateDiscountLine(orderID string, discount *DiscountLine) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "discount_lines", discount.ID),
		method:   http.MethodPut,
		data:     discount,
	})
	if err != nil {
		return err
	}
	return nil
}

// Deletes an existing Discount Line
// https://developers.conekta.com/api?language=bash#delete-discount-line
func (oc *ordersClient) DeleteDiscountLine(orderID, discountID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "discount_lines", discountID),
		method:   http.MethodDelete,
		data:     discountID,
	})
	return err
}

// Creates a new Tax Line
// https://developers.conekta.com/api?language=bash#create-tax-line
func (oc *ordersClient) CreateTaxLine(orderID string, tax *TaxLine) (string, error) {
	res, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "tax_lines"),
		method:   http.MethodPost,
		data:     tax,
	})
	if err != nil {
		return "", err
	}

	var info = map[string]interface{}{}
	json.Unmarshal(res, &info)
	return info["id"].(string), nil
}

// Updates an existing tax line
// https://developers.conekta.com/api?language=bash#update-tax-line
func (oc *ordersClient) UpdateTaxLine(orderID string, tax *TaxLine) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "tax_lines", tax.ID),
		method:   http.MethodPut,
		data:     tax,
	})
	if err != nil {
		return err
	}
	return nil
}

// Deletes an existing tax line
// https://developers.conekta.com/api?language=bash#delete-tax-line
func (oc *ordersClient) DeleteTaxLine(orderID, taxID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "tax_lines", taxID),
		method:   http.MethodDelete,
		data:     taxID,
	})
	return err
}