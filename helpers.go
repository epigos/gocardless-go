package gocardless

import (
	"strings"
	"time"
)

type (
	// Date an alias of time.Time for parsing json dates in the response
	Date struct {
		time.Time
	}
)

// UnmarshalJSON imeplement Marshaler und Unmarshalere interface
func (d *Date) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02", strInput)

	if err != nil {
		return err
	}
	d.Time = newTime
	return nil
}

// Centify amount in floats by multiplying by 100, so 12.25 -> 1225.
// Use when creating payments as amount should be in Pence or Cents
func Centify(amount float64) int {
	return int(amount * 100)
}
