package gocardless

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const (
	rateLimitHeader          = `RateLimit-Limit`
	rateLimitRemainingHeader = `RateLimit-Remaining`
	rateLimitResetHeader     = `RateLimit-Reset`
)

// Response response from the API request, providing access
// to the status code, headers, and body
type Response struct {
	*http.Response
}

// Meta contains pagination cursor for list endpoints
type Meta struct {
	Cursors Cursor `json:"cursors"`
	// Limit Upper bound for the number of objects to be returned. Defaults to 50. Maximum of 500
	Limit int `json:"limit"`
}

// Cursor pagination parameters
type Cursor struct {
	// Before ID of the object immediately following the array of objects to be returned
	Before string `json:"before"`
	// After ID of the object immediately preceding the array of objects to be returned
	After string `json:"after"`
}

// newResponse creates a new response
func newResponse(resp *http.Response) *Response {
	return &Response{resp}
}

// RateLimit the rate limit for each request currently,
// this limit stands at 1000 requests per minute
func (resp *Response) RateLimit() int {
	value, err := strconv.Atoi(resp.Header.Get(rateLimitHeader))
	if err != nil {
		return 0
	}
	return value
}

// RateLimitRemaining indicate how many requests are allowed in the current time window
func (resp *Response) RateLimitRemaining() int {
	value, err := strconv.Atoi(resp.Header.Get(rateLimitRemainingHeader))
	if err != nil {
		return 0
	}
	return value
}

// RateReset indicates the time after which the rate limit will reset
func (resp *Response) RateReset() time.Time {
	value, err := time.Parse(time.RFC1123, resp.Header.Get(rateLimitResetHeader))
	if err != nil {
		t := &time.Time{}
		return *t
	}
	return value
}

// bind decodes response and binds it to struct
func (resp *Response) bind(dst interface{}) error {

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errCtn errorContainer

		err := json.NewDecoder(resp.Body).Decode(&errCtn)
		if err != nil {
			return err
		}

		return errCtn.Error
	}

	if dst != nil {
		err := json.NewDecoder(resp.Body).Decode(dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Meta) String() string {
	bs, _ := json.Marshal(m)
	return string(bs)
}
