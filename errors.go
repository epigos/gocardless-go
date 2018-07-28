package gocardless

import (
	"encoding/json"
)

const (
	// InvalidMethodError details when a request is passed, but the method is invalid in the current context
	InvalidMethodError = `The request Method is invalid`
)

type errorContainer struct {
	Error *Error `json:"error"`
}

// Error base exception class for GoCardless API errors.
// API errors will result in of this
type Error struct {
	DocumentationURL string         `json:"documentation_url"`
	Message          string         `json:"message"`
	RequestID        string         `json:"request_id"`
	Details          []*ErrorDetail `json:"errors"`
	Type             string         `json:"type"`
	Code             int            `json:"code"`
}

func (err Error) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

// ErrorDetail a struct containing the reason for the errors
type ErrorDetail struct {
	Message        string `json:"message"`
	Field          string `json:"field"`
	RequestPointer string `json:"request_pointer"`
}

// RateLimitedExceededError rate limit error
type RateLimitedExceededError struct {
}

func (err *RateLimitedExceededError) Error() string {
	return `Rate Limit exceeded`
}

// InvalidEnvironment invalid environment exception
type InvalidEnvironment error
