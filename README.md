# Go client for Gocardless Pro API

[![Godoc](http://godoc.org/github.com/epigos/gocardless-go?status.svg)](http://godoc.org/github.com/epigos/gocardless-go)
[![Build Status](https://travis-ci.org/epigos/gocardless-go.svg?branch=master)](https://travis-ci.org/epigos/gocardless-go)

This package allows integrating your Golang application with [Gocardless Pro](https://gocardless.com/)

## Installation

Standard go get:

    go get github.com/epigos/gocardless-go

## Coverage

 - Customers
 - Customer Bank Accounts
 - Mandates
 - Payments


 ## Usage

Create a Client instance, providing your access token and the environment you want to use:

```go
package main

import (
    "fmt"
    "os"
    gocardless "github.com/epigos/gocardless-go"
)

func main() {
    token := os.Getenv("GOCARDLESS_ACCESS_TOKEN")
    client := gocardless.NewClient(token, gocardless.SandboxEnvironment)
    
    // get customers
    res, err := client.GetCustomers()
    for _, c := range res.Customers {
        fmt.Println(c)
    }
}
```
## Documentation

- For full usage and examples see the [Godoc](http://godoc.org/github.com/epigos/gocardless-go)
- [Gocardless Pro API reference](https://developer.gocardless.com/api-reference/#overview-getting-started)