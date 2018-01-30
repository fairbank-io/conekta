package conekta

import (
	"errors"
	"testing"
)

func TestConektaClient(t *testing.T) {
	// API key is required
	_, err := New("", nil)
	if err == nil {
		t.Error("failed to detect missing API key")
	}
	
	// Use the test key provided in the public documentation
	client, _ := New("key_eYvWV7gSDkNYXsmr", nil)

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
}
