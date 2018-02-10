package conekta

import (
	"encoding/json"
	"net/http"
	"path"
)

// Defines the public interface required to access available 'orders' methods
type OrdersAPI interface {
	// Creates a new order
	// https://developers.conekta.com/api?language=bash#create-order
	Create(order *Order) error

	// Updates a existing order
	// https://developers.conekta.com/api?language=bash#update-order
	Update(order *Order) error

	// Process a pre-authorized order
	// https://developers.conekta.com/api?language=bash#capture-order
	Capture(orderID string) error

	// A Refund details the amount and reason why an order was refunded
	// https://developers.conekta.com/api?language=bash#refund-order
	Refund(orderID string, r *Refund) error

	// Create a new line item
	// https://developers.conekta.com/api?language=bash#create-line-item
	CreateLineItem(orderID string, item *LineItem) (string, error)

	// Updates a line item
	// https://developers.conekta.com/api?language=bash#update-line-item
	UpdateLineItem(orderID string, item *LineItem) error

	// Deletes a line item
	// https://developers.conekta.com/api?language=bash#delete-line-item
	DeleteLineItem(orderID, itemID string) error

	// Creates a new Discount Line
	// https://developers.conekta.com/api?language=bash#create-discount-line
	CreateDiscountLine(orderID string, discount *DiscountLine) (string, error)

	// Updates an existing Discount Line
	// https://developers.conekta.com/api?language=bash#update-discount-line
	UpdateDiscountLine(orderID string, discount *DiscountLine) error

	// Deletes an existing Discount Line
	// https://developers.conekta.com/api?language=bash#delete-discount-line
	DeleteDiscountLine(orderID, discountID string) error

	// Creates a new Tax Line
	// https://developers.conekta.com/api?language=bash#create-tax-line
	CreateTaxLine(orderID string, tax *TaxLine) (string, error)

	// Updates an existing tax line
	// https://developers.conekta.com/api?language=bash#update-tax-line
	UpdateTaxLine(orderID string, tax *TaxLine) error

	// Deletes an existing tax line
	// https://developers.conekta.com/api?language=bash#delete-tax-line
	DeleteTaxLine(orderID, taxID string) error

	// Creates a new Shipping Line for an existing order
	// https://developers.conekta.com/api?language=bash#create-shipping-line
	CreateShippingLine(orderID string, line *ShippingLine) (string, error)

	// Updates an existing Shipping Line for an existing order
	// https://developers.conekta.com/api?language=bash#update-shipping-line
	UpdateShippingLine(orderID string, line *ShippingLine) error

	// Deletes an existing Shipping Line for an existing order
	// https://developers.conekta.com/api?language=bash#delete-shipping-line
	DeleteShippingLine(orderID, lineID string) error
}

type ordersClient struct {
	c *Client
}

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

func (oc *ordersClient) DeleteLineItem(orderID, itemID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "line_items", itemID),
		method:   http.MethodDelete,
		data:     itemID,
	})
	return err
}

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

func (oc *ordersClient) DeleteDiscountLine(orderID, discountID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "discount_lines", discountID),
		method:   http.MethodDelete,
		data:     discountID,
	})
	return err
}

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

func (oc *ordersClient) DeleteTaxLine(orderID, taxID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "tax_lines", taxID),
		method:   http.MethodDelete,
		data:     taxID,
	})
	return err
}

func (oc *ordersClient) CreateShippingLine(orderID string, line *ShippingLine) (string, error) {
	res, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "shipping_lines"),
		method:   http.MethodPost,
		data:     line,
	})
	if err != nil {
		return "", err
	}

	var info = map[string]interface{}{}
	json.Unmarshal(res, &info)
	return info["id"].(string), nil
}

func (oc *ordersClient) UpdateShippingLine(orderID string, line *ShippingLine) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "shipping_lines", line.ID),
		method:   http.MethodPut,
		data:     line,
	})
	if err != nil {
		return err
	}
	return nil
}

func (oc *ordersClient) DeleteShippingLine(orderID, lineID string) error {
	_, err := oc.c.request(&requestOptions{
		endpoint: baseUrl + path.Join("orders", orderID, "shipping_lines", lineID),
		method:   http.MethodDelete,
		data:     lineID,
	})
	return err
}
