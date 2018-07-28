package gocardless

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	customerEndpoint = "customers"
)

type (
	// Customer struct hold the contact details for a customer
	Customer struct {
		// ID is a unique identifier, beginning with “CU”.
		ID string `json:"id,omitempty"`
		// AddressLine1 is the first line of the customer’s address.
		AddressLine1 string `json:"address_line1"`
		// AddressLine2 is the first line of the customer’s address.
		AddressLine2 string `json:"address_line2"`
		// AddressLine3 is the first line of the customer’s address.
		AddressLint3 string `json:"address_line3"`
		// City is the city of the customer’s address.
		City string `json:"city"`
		// CompanyName is the customer’s company name. Required unless a given_name and family_name are provided.
		CompanyName string `json:"company_name"`
		// CountryCode is the ISO 3166-1 alpha-2 code.
		CountryCode string `json:"country_code"`
		// CreatedAt is a fixed timestamp, recording when the customer was created.
		CreatedAt *time.Time `json:"created_at,omitempty"`
		// Email is the customer's email address
		Email string `json:"email,omitempty"`
		// FamilyName is the customer's surname. Required unless a CompanyName is provided
		FamilyName string `json:"family_name"`
		// GivenName is the customer's first name. Required unless a CompanyName is provided
		GivenName string `json:"given_name"`
		// Language is a ISO 639-1 code. Used as the language for notification emails sent by GoCardless if your
		// organisation does not send its own (see compliance requirements). Currently only “en”, “fr”, “de”, “pt”,
		// “es”, “it”, “nl”, “sv” are supported. If this is not provided, the language will be chosen based on the
		// country_code (if supplied) or default to “en”.
		Language string `json:"language,omitempty"`
		// Metadata is a key-value store of custom data. Up to 3 keys are permitted, with key names up to 50
		// characters and values up to 500 characters.
		Metadata map[string]string `json:"metadata,omitempty"`
		// PostalCode is the customers postal code
		PostalCode string `json:"postal_code"`
		// Region is the customer's address region, county or department
		Region string `json:"region"`
		// SwedishIdentityNumber is for Swedish customers only. The civic/company number (personnummer,
		// samordningsnummer, or organisationsnummer) of the customer. Must be supplied if the customer’s bank
		// account is denominated in Swedish krona (SEK). This field cannot be changed once it has been set.
		SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
	}

	// customerWrapper is a utility struct used to wrap and unwrap the JSON request being passed to the remote API
	customerWrapper struct {
		Customer *Customer `json:"customers"`
	}

	// CustomerListResponse a List response of Customer instances
	CustomerListResponse struct {
		Customers []*Customer `json:"customers"`
		Meta      Meta        `json:"meta,omitempty"`
	}
)

func (cm *Customer) String() string {
	bs, _ := json.Marshal(cm)
	return string(bs)
}

// NewCustomer instantiate a new customer object
func NewCustomer(email, givenName, familyName, line1, line2, city, postalCode, countryCode string) *Customer {
	return &Customer{
		Email:        email,
		GivenName:    givenName,
		FamilyName:   familyName,
		AddressLine1: line1,
		AddressLine2: line2,
		City:         city,
		PostalCode:   postalCode,
		CountryCode:  countryCode,
	}
}

// AddMetadata adds new metadata item to customer object
func (cm *Customer) AddMetadata(key, value string) {
	cm.Metadata[key] = value
}

// CreateCustomer creates a new customer object
//
// Relative endpoint: POST /customers
func (c *Client) CreateCustomer(customer *Customer) error {
	customerReq := &customerWrapper{customer}

	err := c.post(customerEndpoint, customerReq, customerReq)
	if err != nil {
		return err
	}

	return err
}

// GetCustomers returns a cursor-paginated list of your customers.
//
// Relative endpoint: GET /customers
func (c *Client) GetCustomers() (*CustomerListResponse, error) {
	list := &CustomerListResponse{}

	err := c.get(customerEndpoint, list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// GetCustomer retrieves the details of an existing customer.
//
// Relative endpoint: GET /customers/CU123
func (c *Client) GetCustomer(id string) (*Customer, error) {
	wrapper := &customerWrapper{}

	err := c.get(fmt.Sprintf(`%s/%s`, customerEndpoint, id), wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Customer, err
}

// UpdateCustomer Updates a customer object. Supports all of the fields supported when creating a customer.
//
// Relative endpoint: PUT /customers/CU123
func (c *Client) UpdateCustomer(customer *Customer) error {
	id := customer.ID
	// remove unpermitted keys before update
	customer.ID = ""
	customer.CreatedAt = nil

	customerReq := &customerWrapper{customer}

	err := c.put(fmt.Sprintf(`%s/%s`, customerEndpoint, id), customerReq, customerReq)
	if err != nil {
		return err
	}
	return err
}
