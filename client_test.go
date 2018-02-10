package conekta

import (
	"errors"
	"testing"
)

func TestConektaClient(t *testing.T) {
	// API key is required
	_, err := NewClient("", nil)
	if err == nil {
		t.Error("failed to detect missing API key")
	}
	
	// Use the test key provided in the public documentation
	client, _ := NewClient("key_eYvWV7gSDkNYXsmr", nil)

	t.Run("Plans", func(t *testing.T) {
		testPlan := &Plan{
			Name:            "test-plan",
			Currency:        "mxn",
			Amount:          5000,
			Livemode:        false,
			Object:          "plan",
			Interval:        "month",
			Frequency:       1,
			TrialPeriodDays: 10,
			ExpiryCount:     6,
		}

		t.Run("Create", func(t *testing.T) {
			err := client.Plans.Create(testPlan)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Update", func(t *testing.T) {
			_, err := client.Plans.Update(&PlanUpdate{
				ID:     testPlan.ID,
				Name:   "super-new-test-name",
				Amount: 7500,
			})
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Delete", func(t *testing.T) {
			err := client.Plans.Delete(testPlan.ID)
			if err != nil {
				t.Error(err)
			}
		})
	})

	t.Run("Customers", func(t *testing.T) {
		testCustomer := &Customer{
			Name:      "jose",
			Phone:     "+5215542537676",
			Corporate: false,
			Email:     "jose@mail.com",
		}

		t.Run("Create", func(t *testing.T) {
			err := client.Customers.Create(testCustomer)
			if err != nil {
				t.Error(err.(*APIError).Details[0].DebugMessage)
			}
		})

		t.Run("Update", func(t *testing.T) {
			testCustomer.Email = "nuevo@mail.com"
			testCustomer.Phone = "+12026213174"
			if err := client.Customers.Update(testCustomer); err != nil {
				t.Error(err.(*APIError).Details[0].DebugMessage)
			}
		})

		t.Run("PaymentSource", func(t *testing.T) {
			t.Run("Create", func(t *testing.T) {
				err := client.Customers.CreatePaymentSource(testCustomer.ID, "tok_foobar123")
				if err == nil {
					t.Error(errors.New("failed to detect invalid token"))
				}
			})

			t.Run("Update", func(t *testing.T) {
				up := &PaymentSourceUpdate{
					ID:       "pay_invalid_id",
					Name:     "Juanito Perez",
					ExpMonth: "09",
					ExpYear:  "18",
					Address: Address{
						Country:    "US",
						PostalCode: "22202",
					},
				}
				err := client.Customers.UpdatePaymentSource(testCustomer.ID, up)
				if err == nil {
					t.Error(errors.New("failed to detect invalid payment source id"))
				}
			})

			t.Run("Delete", func(t *testing.T) {
				err := client.Customers.DeletePaymentSource(testCustomer.ID, "pay_invalid_id")
				if err == nil {
					t.Error(errors.New("failed to detect invalid payment source id"))
				}
			})
		})

		t.Run("ShippingContact", func(t *testing.T) {
			contact := &ShippingContact{
				Phone: "+5215544332211",
				Address: Address{
					Street1:    "calle 6 910",
					PostalCode: "94510",
					Country:    "MX",
					State:      "Veracruz",
					City:       "Cordoba",
				},
			}

			t.Run("Create", func(t *testing.T) {
				err := client.Customers.CreateShippingContact(testCustomer.ID, contact)
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Update", func(t *testing.T) {
				contact.Receiver = "Laurita Yeye"
				err := client.Customers.UpdateShippingContact(testCustomer.ID, contact)
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Delete", func(t *testing.T) {
				if err := client.Customers.DeleteShippingContact(testCustomer.ID, contact.ID); err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})
		})

		t.Run("Delete", func(t *testing.T) {
			if err := client.Customers.Delete(testCustomer.ID); err != nil {
				t.Error(err.(*APIError).Details[0].DebugMessage)
			}
		})
	})
	
	t.Run("Oders", func(t *testing.T) {
		// Create temporary test customer
		testCustomer := &Customer{
			Name:      "jose",
			Phone:     "+5215542537676",
			Corporate: false,
			Email:     "jose@mail.com",
		}
		if err := client.Customers.Create(testCustomer); err != nil {
			t.Error(err.(*APIError).Details[0].DebugMessage)
		}
		defer client.Customers.Delete(testCustomer.ID)

		// Sample temporary order
		testOrder := &Order{
			Object: "order",
			Currency: "MXN",
			CustomerInfo: CustomerInfo{
				CustomerID: testCustomer.ID,
			},
			LineItems: []LineItem{
				{
					Name: "test digital item",
					Quantity: 1,
					UnitPrice: 5000,
				},
			},
			Charges: []Charge{
				{
					Object: "charge",
					Currency: "MXN",
					PaymentMethod: Card{
						Object: "payment_source",
						Type: "card",
						ExpMonth: "09",
						ExpYear: "19",
						Number: "4242424242424242",
						Name: "Rick Sanchez",
					},
				},
			},
			ShippingContact: ShippingContact{
				Address: Address{
					Street1:    "calle 6 910",
					PostalCode: "94510",
					Country:    "MX",
					State:      "Veracruz",
					City:       "Cordoba",
				},
			},
		}

		t.Run("Create", func(t *testing.T) {
			err := client.Orders.Create(testOrder)
			if err != nil {
				t.Error(err.(*APIError).Details[0].DebugMessage)
			}
		})

		t.Run("Update", func(t *testing.T) {
			testOrder.DiscountLines = []DiscountLine{
				{
					Amount: 1000,
					Type: "campaign",
				},
			}
			err := client.Orders.Update(testOrder)
			if err == nil {
				t.Error("order should not be able to be updated")
			}
		})

		t.Run("Capture", func(t *testing.T) {
			err := client.Orders.Capture(testOrder.ID)
			if err == nil {
				t.Error("order should not be able to be captured")
			}
		})

		t.Run("Refund", func(t *testing.T) {
			err := client.Orders.Refund(testOrder.ID, &Refund{Reason:"other"})
			if err != nil {
				t.Error(err.(*APIError).Details[0].DebugMessage)
			}
		})
		
		t.Run("LineItem", func(t *testing.T) {
			itemID := ""
			t.Run("Create", func(t *testing.T) {
				itemID, err = client.Orders.CreateLineItem(testOrder.ID, &LineItem{
					Name: "another dummy item",
					Quantity: 1,
					UnitPrice: 2000,
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Update", func(t *testing.T) {
				err = client.Orders.UpdateLineItem(testOrder.ID, &LineItem{
					ID: itemID,
					Quantity: 2,
					UnitPrice: 2000,
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Delete", func(t *testing.T) {
				err := client.Orders.DeleteLineItem(testOrder.ID, itemID)
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})
		})

		t.Run("DiscountLine", func(t *testing.T) {
			discountID := ""
			t.Run("Create", func(t *testing.T) {
				discountID, err = client.Orders.CreateDiscountLine(testOrder.ID, &DiscountLine{
					Amount: 1000,
					Type: "coupon",
					Code: "foo-bar",
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Update", func(t *testing.T) {
				err = client.Orders.UpdateDiscountLine(testOrder.ID, &DiscountLine{
					ID: discountID,
					Amount: 1500,
					Type: "coupon",
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Delete", func(t *testing.T) {
				err := client.Orders.DeleteDiscountLine(testOrder.ID, discountID)
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})
		})
		
		t.Run("TaxLine", func(t *testing.T) {
			taxID := ""
			t.Run("Create", func(t *testing.T) {
				taxID, err = client.Orders.CreateTaxLine(testOrder.ID, &TaxLine{
					Amount: 150,
					Description: "IVA",
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Update", func(t *testing.T) {
				err = client.Orders.UpdateTaxLine(testOrder.ID, &TaxLine{
					ID: taxID,
					Amount: 160,
					Description: "IVA",
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Delete", func(t *testing.T) {
				err := client.Orders.DeleteTaxLine(testOrder.ID, taxID)
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})
		})

		t.Run("ShippingLine", func(t *testing.T) {
			shippingID := ""
			t.Run("Create", func(t *testing.T) {
				shippingID, err = client.Orders.CreateShippingLine(testOrder.ID, &ShippingLine{
					Amount: 150,
					Carrier: "UPS",
					TrackingNumber: "foo-bar-123",
					Method: "ground",
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Update", func(t *testing.T) {
				err = client.Orders.UpdateShippingLine(testOrder.ID, &ShippingLine{
					ID: shippingID,
					Method: "air",
				})
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})

			t.Run("Delete", func(t *testing.T) {
				err := client.Orders.DeleteShippingLine(testOrder.ID, shippingID)
				if err != nil {
					t.Error(err.(*APIError).Details[0].DebugMessage)
				}
			})
		})
	})
}
