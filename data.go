package conekta

// https://developers.conekta.com/api?language=bash#shipping-contact
type Address struct {
	// The first line for the shipping address. Usually used for the street and the number
	Street1 string `json:"street1,omitempty"`

	// The second line for the shipping address. Usually used for the internal number,
	// suite, residential complex or county
	Street2 string `json:"street2,omitempty"`

	// The city for the address
	City string `json:"city,omitempty"`

	// The state for the address
	State string `json:"state,omitempty"`

	// Use the two-digit format ISO 3166-1
	Country string `json:"country,omitempty"`

	// The postal code for the address
	PostalCode string `json:"postal_code,omitempty"`

	// Boolean value that indicates whether it is a residential address
	// By default is taken as true (Optional)
	Residential bool `json:"residential"`
}

// Describes the customer's information for the shipment.
// https://developers.conekta.com/api?language=bash#shipping-contact
type ShippingContact struct {
	// Unique identifier, assigned at random
	ID string `json:"id,omitempty"`

	// Contact's phone number
	Phone string `json:"phone,omitempty"`

	// Contact's name (optional)
	Receiver string `json:"receiver,omitempty"`

	// Adjacent streets to the shipping address (optional)
	BetweenStreets string `json:"between_streets,omitempty"`

	// Shipping address
	Address Address `json:"address,omitempty"`
}

// Represents the line items in the order.
// https://developers.conekta.com/api?language=bash#line-item
type LineItem struct {
	// Unique identifier for the line item
	ID string `json:"id,omitempty"`

	// Line item's name
	Name string `json:"name,omitempty"`

	// Short description of the item (<250). (optional)
	Description string `json:"description,omitempty"`

	// Line item's price in cents
	UnitPrice uint32 `json:"unit_price,omitempty"`

	// Line item's quantity for the order
	Quantity uint32 `json:"quantity,omitempty"`

	// The Storage Keeping Unit designated by the company (optional)
	SKU string `json:"sku,omitempty"`

	// Array containing the line item's categories (optional)
	Tags []string `json:"tags,omitempty"`

	// Line item's brand (optional)
	Brand string `json:"brand,omitempty"`

	// Map containing additional information related to the line item (optional)
	Metadata map[string]string `json:"metadata,omitempty"`
}

// A Shipping Line describes the shipment details for an order such as the method,
// the amount, the carrier and the tracking number.
// https://developers.conekta.com/api?language=bash#shipping-line
type ShippingLine struct {
	// Unique identifier, assigned at random
	ID string `json:"id,omitempty"`

	// The shipping cost, in cents
	Amount uint32 `json:"amount,omitempty"`

	// Tracking number provided by the carrier (optional)
	TrackingNumber string `json:"tracking_number,omitempty"`

	// Name of the carrier
	Carrier string `json:"carrier,omitempty"`

	// Shipping method (optional)
	Method string `json:"method,omitempty"`

	// Map containing additional information related to the shipping line (optional)
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Describes the taxes of the order.
// https://developers.conekta.com/api?language=bash#tax-line
type TaxLine struct {
	// Unique identifier for the tax
	ID string `json:"id,omitempty"`

	// Tax code
	Description string `json:"description,omitempty"`

	// The tax amount to be paid
	Amount uint32 `json:"amount,omitempty"`

	// Map containing additional information related to the tax line (optional)
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Describes the discounts to be applied to the order
// https://developers.conekta.com/api?language=bash#discount-line
type DiscountLine struct {
	// Unique identifier, assigned at random
	ID string `json:"id,omitempty"`

	// Discount line's code
	Code string `json:"code,omitempty"`

	// It can be loyalty, campaign, coupon or sign
	Type string `json:"type,omitempty"`

	// The discount's amount, in cents
	Amount uint32 `json:"amount,omitempty"`
}

// Map containing information about the order's customer.
type CustomerInfo struct {
	// Id of the customer. This is required if the other customer_info fields are not sent
	CustomerID string `json:"customer_id,omitempty"`

	// Customer's name (optional if customer_id is sent)
	Name string `json:"name,omitempty"`

	// Customer's phone (optional if customer_id is sent)
	Phone string `json:"phone,omitempty"`

	// Customer's email (optional if customer_id is sent)
	Email string `json:"email,omitempty"`

	// It states if the customer is corporate or not, the default value is false (optional)
	Corporate bool `json:"corporate"`
}

// The Charge object contains all information related to the payment of an order.
// This object belongs to the order object.
// https://developers.conekta.com/api?language=bash#charge
type Charge struct {
	// Unique identifier, assigned at random
	ID string `json:"id,omitempty"`

	// Object class. In this case, "charge"
	Object string `json:"object,omitempty"`

	// Status of the charge
	Status string `json:"status,omitempty"`

	// Id of the order that the charge belongs to
	OrderID string `json:"order_id,omitempty"`

	// Payment method of the charge
	PaymentMethod Card `json:"payment_method,omitempty"`

	// Date of the charge's creation
	CreatedAt uint32 `json:"created_at,omitempty"`

	// Currency of the charge. A 3-letter code of the International Standard ISO 4217
	Currency string `json:"currency,omitempty"`

	// The charge's amount, in cents
	Amount uint32 `json:"amount,omitempty"`

	// Amount of the fee in cents
	Fee uint32 `json:"fee,omitempty"`

	// Monthly installments in which the charge will be divided, with no interest added.
	// Conekta offers monthly installments of 3, 6, 9 and 12 payments.
	MonthlyInstallments uint32 `json:"monthly_installments,omitempty"`

	// false: Sandbox Mode. true: Production Mode
	Livemode bool `json:"livemode"`
}

// A refund details the amount and reason why an order was refunded
// https://developers.conekta.com/api?language=bash#refund-order
type Refund struct {
	// Order's id
	ID string `json:"id,omitempty"`

	// Reason for refund
	// requested_by_client
	// cannot_be_fulfilled
	// duplicated_transaction
	// suspected_fraud
	// other
	Reason string `json:"reason,omitempty"`

	// If you want to partially refund and order
	Amount uint32 `json:"amount,omitempty"`
}

// An Order represents a purchase. It contains all the details related to it, including
// payment method, shipment, charges, discounts, taxes and the products.
// https://developers.conekta.com/api?language=bash#order
type Order struct {
	// Unique identifier assigned at random
	ID string `json:"id,omitempty"`

	// Object class. In this case, "order"
	Object string `json:"object,omitempty"`

	// Date when the order was created
	CreatedAt uint32 `json:"created_at,omitempty"`

	// Date when the order was last updated
	UpdatedAt uint32 `json:"updated_at,omitempty"`

	// Currency of the charge, ISO 4217
	Currency string `json:"currency,omitempty"`

	// List of the products being sold in the order. It must contain at least one product
	LineItems []LineItem `json:"line_items,omitempty"`

	// List of shipping costs. This parameter is optional
	ShippingLines []ShippingLine `json:"shipping_lines,omitempty"`

	// List of the taxes
	TaxLines []TaxLine `json:"tax_lines,omitempty"`

	// List of the discounts applied to the order.
	DiscountLines []DiscountLine `json:"discount_lines,omitempty"`

	// false: Sandbox Mode. true: Production Mode
	Livemode bool `json:"livemode"`

	// Map containing additional information related to the order
	Metadata map[string]string `json:"metadata,omitempty"`

	// Mandatory when a shipping_line is included in the order. If the order is sent without a
	// shipping_contact, the customer's default shipping_contact will be used
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`

	// Amount calculated based on line_items, shipping_lines, tax_lines and discount_lines
	Amount uint32 `json:"amount,omitempty"`

	// Amount refunded through a call to: orders/:order_id/refund
	AmountRefunded uint32 `json:"amount_refunded,omitempty"`

	// Status of the order's payment. This field is set by the system, it can be:
	// payment_pending, declined, expired, paid, refunded, partially_refunded, charged_back, pre_authorized and voided
	PaymentStatus string `json:"payment_status,omitempty"`

	// Information about the order's customer
	CustomerInfo CustomerInfo `json:"customer_info,omitempty"`

	// List of the charges generated to cover the order amount
	Charges []Charge `json:"charges,omitempty"`

	// States if the charges of the order should be preauthorized
	PreAuthorize bool `json:"pre_authorize"`
}

// The Payment Source object describes a payment method. This can be online (card payments)
// or offline OXXO and SPEI. Remember that for offline payments you will need to add
// a webhook listener
// https://developers.conekta.com/api?language=bash#payment-source
type PaymentSource struct {
	// Unique identifier randomly assigned
	ID string `json:"id,omitempty"`

	// Object's class. In this case "payment_source"
	Object string `json:"object,omitempty"`

	// Payment source's type. At the time the only type supported is "card"
	Type string `json:"type,omitempty"`

	// Date when the payment source was created
	CreatedAt uint32 `json:"created_at,omitempty"`

	// Last 4 digits of the card
	Last4 string `json:"last4,omitempty"`

	// Name of the card holder
	Name string `json:"name,omitempty"`

	// Expiration month of the card
	ExpMonth string `json:"exp_month,omitempty"`

	// Expiration year of the card
	ExpYear string `json:"exp_year,omitempty"`

	// Card's brand
	Brand string `json:"brand,omitempty"`

	// Id of the customer that owns the payment source
	ParentID string `json:"parent_id,omitempty"`
}

// Updates an existing payment source
// https://developers.conekta.com/api?language=bash#update-payment-source
type PaymentSourceUpdate struct {
	// Id of the Payment Source
	ID string `json:"id,omitempty,omitempty"`

	// Name of the card holder
	Name string `json:"name,omitempty,omitempty"`

	// Expiration month of the card
	ExpMonth string `json:"exp_month,omitempty,omitempty"`

	// Expiration year of the card
	ExpYear string `json:"exp_year,omitempty,omitempty"`

	// Address of the cardholder
	Address Address `json:"address,omitempty,omitempty"`
}

// Subscriptions bill your client a fixed amount on a recurring basis. You can change the
// plan of a subscription, pause, cancel and resume a subscription as you wish
// https://developers.conekta.com/api?language=bash#subscription
type Subscription struct {
	// Unique identifier
	ID string `json:"id,omitempty"`

	// Object class. For this model, "subscription"
	Object string `json:"object,omitempty"`

	// Date of the subscription creation
	CreatedAt uint32 `json:"created_at,omitempty"`

	// Date of the subscription cancelation
	CanceledAt uint32 `json:"canceled_at,omitempty"`

	// Date of the subscription pause
	PausedAt uint32 `json:"paused_at,omitempty"`

	// Date of the billing cycle start
	BillingCycleStart uint32 `json:"billing_cycle_start,omitempty"`

	// Date of the billing cycle end
	BillingCycleEnd uint32 `json:"billing_cycle_end,omitempty"`

	// Date of the trial start
	TrialStart uint32 `json:"trial_start,omitempty"`

	// Date of the trial end
	TrialEnd uint32 `json:"trial_end,omitempty"`

	// Id of the plan assigned to the subscription
	PlanID string `json:"plan_id,omitempty"`

	// Status of the subscription. Allowed values are:
	// in_trial, active, past_due, paused, and canceled
	Status string `json:"status,omitempty"`
}

// Customers allow you to store payment methods for clients and set up subscriptions
// https://developers.conekta.com/api?language=bash#customer
type Customer struct {
	// Customer's unique identifier
	ID string `json:"id,omitempty"`

	// Customer's name
	Name string `json:"name,omitempty"`

	// Customer's phone number (international format)
	Phone string `json:"phone,omitempty"`

	// Customer's email
	Email string `json:"email,omitempty"`

	// Plan secondary id
	PlanID string `json:"plan_id,omitempty"`

	// Indicates whether a user is corporate or not
	Corporate bool `json:"corporate"`

	// Payment sources available
	PaymentSources []PaymentSource `json:"payment_sources,omitempty"`

	// Shipping contacts available
	ShippingContacts []ShippingContact `json:"shipping_contacts,omitempty"`

	// Subscriptions bill your client a fixed amount on a recurring basis.
	// You can change the plan of a subscription, pause, cancel and resume a subscription
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// Plans are templates for subscriptions. They allow you to define the amount and
// frequency with which you would like to bill your clients.
// https://developers.conekta.com/api?language=bash#plan
type Plan struct {
	// Unique identifier assigned at random
	ID string `json:"id,omitempty"`

	// Object class. In this case, "plan"
	Object string `json:"object,omitempty"`

	// Date when the order was created
	CreatedAt uint32 `json:"created_at,omitempty"`

	// false: Sandbox Mode. true: Production Mode
	Livemode bool `json:"livemode"`

	// Plan's name
	Name string `json:"name,omitempty"`

	// Charge's amount in cents
	Amount uint32 `json:"amount,omitempty"`

	// Currency of the charge. A 3-letter code of the International Standard ISO 4217
	Currency string `json:"currency,omitempty"`

	// The interval for the charge. For example, to charge a customer every 2 months,
	// set the interval attribute to month and the frequency to 2
	Interval string `json:"interval,omitempty"`

	// The frequency for the charge. For example, to charge a customer every 2 months,
	// set the interval attribute to month and the frequency to 2
	Frequency uint32 `json:"frequency,omitempty"`

	// Days of the trial's duration
	TrialPeriodDays uint32 `json:"trial_period_days,omitempty"`

	// Number of charges that will be made before the subscription expires
	ExpiryCount uint32 `json:"expiry_count,omitempty"`
}

// Updates plan data
// https://developers.conekta.com/api?language=bash#update-plan
type PlanUpdate struct {
	// Unique plan identifier. Remember you can't change it
	ID string `json:"id,omitempty"`

	// Plan's name
	Name string `json:"name,omitempty"`

	// Charge's amount in cents
	Amount uint32 `json:"amount,omitempty"`
}

// Card enable to charge orders directly to a user plastic card
// PaymentSource available documentation is incomplete
// https://developers.conekta.com/api?language=bash#payment-source
type Card struct {
	// Object's class. In this case "payment_source"
	Object string `json:"object,omitempty"`

	// Payment source's type. At the time the only type supported is "card"
	Type string `json:"type,omitempty"`

	// Number of the card
	Number string `json:"number,omitempty"`

	// Expiration month of the card
	ExpMonth string `json:"exp_month,omitempty"`

	// Expiration year of the card
	ExpYear string `json:"exp_year,omitempty"`

	// Card's brand
	Brand string `json:"brand,omitempty"`

	// Card's holder name
	Name string `json:"name,omitempty"`
}