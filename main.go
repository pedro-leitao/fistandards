package main

import (
	"fmt"

	"github.com/pedro-leitao/fistandards/iban"
	"github.com/pedro-leitao/fistandards/pan"
)

func main() {

	var iban iban.Iban
	var pan pan.Pan
	var normalized string
	var err error

	if normalized, err = iban.Validate("GB82-WEST 1234 5698 7654 32"); err != nil {
		fmt.Printf("%v: %v\n", normalized, err.Error())
	} else {
		fmt.Printf("%v is a valid IBAN\n", normalized)
		fmt.Printf("BBAN: %v, Country code: %v\n", iban.Bban, iban.CountryCode)
	}

	if normalized, err = pan.Validate("5460.9762.4968.5093"); err != nil {
		fmt.Printf("%v: %v\n", normalized, err.Error())
	} else {
		fmt.Printf("%v is a valid PAN\n", normalized)
		fmt.Printf("Issuer: %v, IIN: %v\n", pan.Issuer, pan.Iin)
	}
}
