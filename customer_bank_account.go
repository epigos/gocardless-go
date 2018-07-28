package gocardless

import (
	"encoding/json"
	"fmt"
)

const (
	bankAccountEndpoint = "customer_bank_accounts"
)

type (
	// CustomerBankAccount Customer Bank Accounts hold the bank details of a customer
	// They always belong to a customer, and may be linked to several Direct Debit mandates.
	CustomerBankAccount struct {
		// ID is a unique identifier, beginning with “CU”.
		ID string `json:"id,omitempty"`
		// AccountHolderName Name of the account holder, as known by the bank.
		// Usually this matches the name of the linked customer.
		// This field will be transliterated, upcased and truncated to 18 characters.
		AccountHolderName string `json:"account_holder_name"`
		// AccountNumber Bank account number. Alternatively you can provide an iban
		AccountNumber string `json:"account_number"`
		// BankCode Bank code
		BankCode string `json:"bank_code,omitempty"`
		// BankCode Bank code
		BranchCode string `json:"branch_code"`
		// AccountNumberEnding Last two digits of account number
		AccountNumberEnding string `json:"account_number_ending,omitempty"`
		// BankName Name of bank, taken from the bank details
		BankName string `json:"bank_name,omitempty"`
		// CountryCode is the ISO 3166-1 alpha-2 code.
		CountryCode string `json:"country_code"`
		// Currency currency code, defaults to national currency of country_code
		Currency string `json:"currency,omitempty"`
		// IBAN International Bank Account Number
		IBAN string `json:"iban,omitempty"`
		// Metadata is a key-value store of custom data. Up to 3 keys are permitted, with key names up to 50
		// characters and values up to 500 characters.
		Metadata map[string]string `json:"metadata,omitempty"`
		// Links links constains customers id
		Links links `json:"links"`
	}
	links struct {
		// CustomerID ID of customer who owns the bank account
		CustomerID string `json:"customer"`
		// CustomerBankAccountToken ID of a customer bank account token to use in place of bank account parameters.
		CustomerBankAccountToken string `json:"customer_bank_account_token,omitempty"`
	}
)

// customerBankAccountWrapper is a utility struct used to wrap and unwrap the JSON request being passed to the remote API
type customerBankAccountWrapper struct {
	CustomerBankAccount *CustomerBankAccount `json:"customer_bank_accounts"`
}

// CustomerBankAccountListResponse a List response of CustomerBankAccount instances
type CustomerBankAccountListResponse struct {
	CustomerBankAccounts []*CustomerBankAccount `json:"customer_bank_accounts"`
	Meta                 Meta                   `json:"meta,omitempty"`
}

func (ca *CustomerBankAccount) String() string {
	bs, _ := json.Marshal(ca)
	return string(bs)
}

// NewCustomerBankAccount instantiate a new customer bank account object
func NewCustomerBankAccount(accountNumber, accountName, branchCode, countryCode, customerID string) *CustomerBankAccount {
	return &CustomerBankAccount{
		AccountNumber:     accountNumber,
		BranchCode:        branchCode,
		AccountHolderName: accountName,
		CountryCode:       countryCode,
		Links:             links{CustomerID: customerID},
	}
}

// AddMetadata adds new metadata item to customer object
func (ca *CustomerBankAccount) AddMetadata(key, value string) {
	ca.Metadata[key] = value
}

// CreateCustomerBankAccount creates a new customer bank account object
// Relative endpoint: POST /customer_bank_accounts
func (c *Client) CreateCustomerBankAccount(cba *CustomerBankAccount) error {
	cbaReq := &customerBankAccountWrapper{cba}

	err := c.post(bankAccountEndpoint, cbaReq, cbaReq)
	if err != nil {
		return err
	}

	return err
}

// GetCustomerBankAccounts returns a cursor-paginated list of your bank accounts
// Relative endpoint: GET /customer_bank_accounts
func (c *Client) GetCustomerBankAccounts() (*CustomerBankAccountListResponse, error) {
	list := &CustomerBankAccountListResponse{}

	err := c.get(bankAccountEndpoint, list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// GetCustomerBankAccount Retrieves the details of an existing bank account
// Relative endpoint: GET /customer_bank_accounts/BA123
func (c *Client) GetCustomerBankAccount(id string) (*CustomerBankAccount, error) {
	wrapper := &customerBankAccountWrapper{}

	err := c.get(fmt.Sprintf(`%s/%s`, bankAccountEndpoint, id), wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.CustomerBankAccount, err
}

// UpdateCustomerBankAccount Updates a customer bank account object. Only the metadata parameter is allowed
// Relative endpoint: PUT /customer_bank_accounts/BA123
func (c *Client) UpdateCustomerBankAccount(id string, cba *CustomerBankAccount) error {
	// remove unpermitted keys before update
	cbaMeta := map[string]interface{}{
		"customer_bank_accounts": map[string]interface{}{
			"metadata": cba.Metadata,
		},
	}
	cbaRes := &customerBankAccountWrapper{cba}

	err := c.put(fmt.Sprintf(`%s/%s`, bankAccountEndpoint, id), cbaMeta, cbaRes)
	if err != nil {
		return err
	}
	return err
}
