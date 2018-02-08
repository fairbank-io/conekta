# Electrum Client

[![Build Status](https://travis-ci.org/fairbank-io/conekta.svg?branch=master)](https://travis-ci.org/fairbank-io/conekta)
[![GoDoc](https://godoc.org/github.com/fairbank-io/conekta?status.svg)](https://godoc.org/github.com/fairbank-io/conekta)
[![Version](https://img.shields.io/github/tag/fairbank-io/conekta.svg)](https://github.com/fairbank-io/conekta/releases)
[![Software License](https://img.shields.io/badge/license-MIT-red.svg)](LICENSE)

Pure Go [Conekta](https://conectka.com) client implementation.

## Example

```go
// Start a new client instance with default options
client, err := conekta.NewClient("API_KEY", nil)

// Create new customer
testCustomer := &Customer{
    Name:      "jose",
    Phone:     "+5215542537676",
    Corporate: false,
    Email:     "jose@mail.com",
}
err := client.Customers.Create(testCustomer)
if err != nil {
	// Handle error
}
```
