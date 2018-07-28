package gocardless

import (
	"fmt"
	"os"
)

func ExampleCustomer() {
	// Create a Client instance, providing your access token and the environment you want to use
	token := os.Getenv("GOCARDLESS_ACCESS_TOKEN")
	client := NewClient(token, SandboxEnvironment)

	// create customer
	cm := NewCustomer("user@example.com", "Frank", "Osborne", "27 Acer Road", "Apt 2", "London", "E8 3GX", "GB")
	cm.AddMetadata("salesforce_id", "ABCD1234")
	err := client.CreateCustomer(cm)

	if err != nil {
		panic(err)
	}
	fmt.Println(cm)

	// retrieve customer
	cm, err = client.GetCustomer(cm.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(cm)
	// get customers
	res, err := client.GetCustomers()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// update customer
	cm.CompanyName = "Google.com"
	err = client.UpdateCustomer(cm)
	if err != nil {
		panic(err)
	}
	fmt.Println(cm)
}
