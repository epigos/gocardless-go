package gocardless

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	paymentEndpoint = "payments"
)

type (
	// Payment objects represent payments from a customer to a creditor, taken against a Direct Debit payment.
	Payment struct {
		// ID is a unique identifier, beginning with "PM".
		ID string `json:"id,omitempty"`
		// Amount in pence (GBP), cents (AUD/EUR), öre (SEK), or øre (DKK).
		// e.g 1000 is 10 GBP in pence
		Amount int `json:"amount"`
		// AmountRefunded is amount refunded in pence/cents/öre/øre.
		AmountRefunded int `json:"amount_refunded,omitempty"`
		// ChargeDate A future date on which the payment should be collected.
		// If not specified, the payment will be collected as soon as possible
		ChargeDate *Date `json:"charge_date,omitempty"`
		// CreatedAt is a fixed timestamp, recording when the payment was created.
		CreatedAt *time.Time `json:"created_at,omitempty"`
		// Currency currency code, defaults to national currency of country_code
		Currency string `json:"currency"`
		// Description A human-readable description of the payment
		Description string `json:"description,omitempty"`
		// Metadata is a key-value store of custom data. Up to 3 keys are permitted, with key names up to 50
		// characters and values up to 500 characters.
		Metadata map[string]string `json:"metadata,omitempty"`
		// Reference An optional payment reference that will appear on your customer’s bank statement
		Reference string `json:"reference,omitempty"`
		// Status status of payment.
		Status string `json:"status,omitempty"`
		// Links to cusomer and payment
		Links paymentLinks `json:"links"`
		// AppFee The amount to be deducted from the payment as the OAuth app’s fee, in pence/cents/öre/øre
		AppFee int `json:"app_fee,omitempty"`
	}
	paymentLinks struct {
		CreditorID     string `json:"creditor,omitempty"`
		PayoutID       string `json:"payout,omitempty"`
		SubscriptionID string `json:"subscription,omitempty"`
		MandateID      string `json:"mandate,omitempty"`
	}
	// paymentWrapper is a utility struct used to wrap and unwrap the JSON request being passed to the remote API
	paymentWrapper struct {
		Payment *Payment `json:"payments"`
	}

	// PaymentListResponse a List response of Payment instances
	PaymentListResponse struct {
		Payments []*Payment `json:"payments"`
		Meta     Meta       `json:"meta,omitempty"`
	}
)

func (p *Payment) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}

// NewPayment instantiate new payment object
func NewPayment(amount int, currency, mandateID string) *Payment {
	return &Payment{
		Amount:   amount,
		Currency: currency,
		Links:    paymentLinks{MandateID: mandateID},
	}
}

// AddMetadata adds new metadata item to payment object
func (p *Payment) AddMetadata(key, value string) {
	p.Metadata[key] = value
}

// CreatePayment creates a new payment object.
//
// Relative endpoint: POST /payments
func (c *Client) CreatePayment(payment *Payment) error {
	paymentReq := &paymentWrapper{payment}

	err := c.post(paymentEndpoint, paymentReq, paymentReq)
	if err != nil {
		return err
	}

	return err
}

// GetPayments returns a cursor-paginated list of your payments.
//
// Relative endpoint: GET /payments
func (c *Client) GetPayments() (*PaymentListResponse, error) {
	list := &PaymentListResponse{}

	err := c.get(paymentEndpoint, list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// GetPayment retrieves the details of an existing payment.
//
// Relative endpoint: GET /payments/PM123
func (c *Client) GetPayment(id string) (*Payment, error) {
	wrapper := &paymentWrapper{}

	err := c.get(fmt.Sprintf(`%s/%s`, paymentEndpoint, id), wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Payment, err
}

// UpdatePayment Updates a payment object. Supports all of the fields supported when creating a payment.
//
// Relative endpoint: PUT /payments/PM123
func (c *Client) UpdatePayment(payment *Payment) error {
	// allows only metadata
	paymentMeta := map[string]interface{}{
		"payments": map[string]interface{}{
			"metadata": payment.Metadata,
		},
	}

	paymentReq := &paymentWrapper{payment}

	err := c.put(fmt.Sprintf(`%s/%s`, paymentEndpoint, payment.ID), paymentMeta, paymentReq)
	if err != nil {
		return err
	}
	return err
}

// CancelPayment immediately cancels a payment and all associated cancellable payments.
//
// Relative endpoint: POST /payments/PM123/actions/cancel
func (c *Client) CancelPayment(payment *Payment) error {
	// allows only metadata
	pMeta := map[string]interface{}{
		"payments": map[string]interface{}{
			"metadata": payment.Metadata,
		},
	}
	wrapper := &paymentWrapper{payment}
	err := c.post(fmt.Sprintf(`%s/%s/actions/cancel`, paymentEndpoint, payment.ID), pMeta, wrapper)
	if err != nil {
		return err
	}
	return err
}

// RetryPayment Retries a failed payment if the underlying mandate is active.
//
// Relative endpoint: POST /payments/PM123/actions/retry
func (c *Client) RetryPayment(payment *Payment) error {
	// allows only metadata
	pMeta := map[string]interface{}{
		"payments": map[string]interface{}{
			"metadata": payment.Metadata,
		},
	}
	wrapper := &paymentWrapper{payment}
	err := c.post(fmt.Sprintf(`%s/%s/actions/retry`, paymentEndpoint, payment.ID), pMeta, wrapper)
	if err != nil {
		return err
	}
	return err
}
