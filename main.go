package main

import (
	"fmt"

	"github.com/pedro-leitao/fistandards/iban"
)

func main() {

	var iban iban.Iban
	var normalized string
	var err error

	if normalized, err = iban.Set("GB82-WEST 1234 5698 7654 32"); err != nil {
		fmt.Printf("%v: $v\n", normalized, err.Error())
	} else {
		fmt.Printf("%v is a valid IBAN\n", normalized)
	}

}
