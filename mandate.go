package gocardless

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	mandateEndpoint = "mandates"
)

type (
	// Mandate Mandates represent the Direct Debit mandate with a customer.
	Mandate struct {
		// ID is a unique identifier, beginning with “MD”.
		ID string `json:"id,omitempty"`
		// CreatedAt is a fixed timestamp, recording when the mandate was created.
		CreatedAt *time.Time `json:"created_at,omitempty"`
		// Metadata is a key-value store of custom data. Up to 3 keys are permitted, with key names up to 50
		// characters and values up to 500 characters.
		Metadata map[string]string `json:"metadata,omitempty"`
		// NextPossibleChargeDate The earliest date a newly created payment for this mandate could be charged
		NextPossibleChargeDate *time.Time `json:"next_possible_charge_date,omitempty"`
		// PaymentRequireApproval Boolean value showing whether payments and
		// subscriptions under this mandate require approval via an automated email before being processed
		PaymentRequireApproval bool `json:"payments_require_approval,omitempty"`
		// Reference Unique reference
		// Different schemes have different length and character set requirements
		// GoCardless will generate a unique reference satisfying the different scheme requirements if this field is left blank
		Reference string `json:"reference,omitempty"`
		// Scheme Direct Debit scheme to which this mandate and associated payments are submitted
		Scheme string `json:"scheme,omitempty"`
		// Status status of mandate.
		Status string `json:"status,omitempty"`
		// Links links to cusomer and bacnk accounts
		Links mandateLinks `json:"links"`
	}
	mandateLinks struct {
		CreditorID            string `json:"creditor,omitempty"`
		CustomerID            string `json:"customer,omitempty"`
		CustomerBankAccountID string `json:"customer_bank_account,omitempty"`
		NewMandateID          string `json:"new_mandate,omitempty"`
	}
	// mandateWrapper is a utility struct used to wrap and unwrap the JSON request being passed to the remote API
	mandateWrapper struct {
		Mandate *Mandate `json:"mandates"`
	}

	// MandateListResponse a List response of Mandate instances
	MandateListResponse struct {
		Mandates []*Mandate `json:"mandates"`
		Meta     Meta       `json:"meta,omitempty"`
	}
)

func (m *Mandate) String() string {
	bs, _ := json.Marshal(m)
	return string(bs)
}

// NewMandate instantiate new mandate object
func NewMandate(bankAccountID string) *Mandate {
	return &Mandate{
		Links: mandateLinks{CustomerBankAccountID: bankAccountID},
	}
}

// AddMetadata adds new metadata item to mandate object
func (m *Mandate) AddMetadata(key, value string) {
	m.Metadata[key] = value
}

// CreateMandate creates a new mandate object.
//
// Relative endpoint: POST /mandates
func (c *Client) CreateMandate(mandate *Mandate) error {
	mandateReq := &mandateWrapper{mandate}

	err := c.post(mandateEndpoint, mandateReq, mandateReq)
	if err != nil {
		return err
	}

	return err
}

// GetMandates returns a cursor-paginated list of your mandates.
//
// Relative endpoint: GET /mandates
func (c *Client) GetMandates() (*MandateListResponse, error) {
	list := &MandateListResponse{}

	err := c.get(mandateEndpoint, list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// GetMandate retrieves the details of an existing mandate.
//
// Relative endpoint: GET /mandates/MD123
func (c *Client) GetMandate(id string) (*Mandate, error) {
	wrapper := &mandateWrapper{}

	err := c.get(fmt.Sprintf(`%s/%s`, mandateEndpoint, id), wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Mandate, err
}

// UpdateMandate Updates a mandate object. Supports all of the fields supported when creating a mandate.
//
// Relative endpoint: PUT /mandates/MD123
func (c *Client) UpdateMandate(mandate *Mandate) error {
	// remove unpermitted keys before update
	cbaMeta := map[string]interface{}{
		"mandates": map[string]interface{}{
			"metadata": mandate.Metadata,
		},
	}

	mandateReq := &mandateWrapper{mandate}

	err := c.put(fmt.Sprintf(`%s/%s`, mandateEndpoint, mandate.ID), cbaMeta, mandateReq)
	if err != nil {
		return err
	}
	return err
}

// CancelMandate immediately cancels a mandate and all associated cancellable payments.
//
// Relative endpoint: POST /mandates/MD123/actions/cancel
func (c *Client) CancelMandate(id string) (*Mandate, error) {
	wrapper := &mandateWrapper{}
	err := c.post(fmt.Sprintf(`%s/%s/actions/cancel`, mandateEndpoint, id), nil, wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Mandate, err
}

// ReinstateMandate Reinstates a cancelled or expired mandate to the banks.
//
// Relative endpoint: POST /mandates/MD123/actions/reinstate
func (c *Client) ReinstateMandate(id string) (*Mandate, error) {
	wrapper := &mandateWrapper{}
	err := c.post(fmt.Sprintf(`%s/%s/actions/reinstate`, mandateEndpoint, id), nil, wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Mandate, err
}
