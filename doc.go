/*
Package gocardless is a Go client for interacting with the GoCardless Pro API.

Example:

Create a Client instance, providing your access token and the environment you want to use:

  package main
  import (
    "fmt"
	"os"
    gocardless "github.com/epigos/gocardless-go"
  )
  func main() {
	token := os.Getenv("GOCARDLESS_ACCESS_TOKEN")
	client := gocardless.NewClient(token, gocardless.Sandbox)

	res, err := client.GetCustomers()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
  }

Learn more about GoCardless Pro API https://developer.gocardless.com/
*/
package gocardless
