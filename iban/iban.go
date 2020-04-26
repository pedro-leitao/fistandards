// Package iban is a set of methods to validate and process IBAN representations.
// See https://www.mobilefish.com/services/bban_iban/bban_iban.php for structure details.
package iban

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

var countryLengths = map[string]int{
	"AD": 24, "AE": 23, "AT": 20, "AZ": 28, "BA": 20, "BE": 16, "BG": 22, "BH": 22, "BR": 29,
	"CH": 21, "CR": 21, "CY": 28, "CZ": 24, "DE": 22, "DK": 18, "DO": 28, "EE": 20, "ES": 24,
	"FI": 18, "FO": 18, "FR": 27, "GB": 22, "GI": 23, "GL": 18, "GR": 27, "GT": 28, "HR": 21,
	"HU": 28, "IE": 22, "IL": 23, "IS": 26, "IT": 27, "JO": 30, "KW": 30, "KZ": 20, "LB": 28,
	"LI": 21, "LT": 20, "LU": 20, "LV": 21, "MC": 27, "MD": 24, "ME": 22, "MK": 19, "MR": 27,
	"MT": 31, "MU": 30, "NL": 18, "NO": 15, "PK": 24, "PL": 28, "PS": 29, "PT": 25, "QA": 29,
	"RO": 24, "RS": 22, "SA": 24, "SE": 24, "SI": 19, "SK": 24, "SM": 27, "TN": 24, "TR": 26,
}

// Iban is the representation of an IBAN value
type Iban struct {
	iban            string
	countryCode     string
	ibanCheckDigits string
	bban            string
}

// Set the object to a given IBAN string, and check its validity returning a normalized
// representation and an error in case verification failed.
func (c *Iban) Set(s string) (string, error) {

	if len(s) < 2 {
		return s, errors.New("Invalid length")
	}

	clean := strings.TrimSpace(s)
	clean = strings.ReplaceAll(clean, ".", "")
	clean = strings.ReplaceAll(clean, "-", "")
	clean = strings.ReplaceAll(clean, " ", "")
	clean = strings.ReplaceAll(clean, "\t", "")
	clean = strings.ToUpper(clean)

	cl, ok := countryLengths[clean[0:2]]
	// Invalid country code
	if !ok {
		return clean, errors.New("Invalid country code")
	}

	// Invalid length
	if len(s) < cl {
		return clean, errors.New("Invalid length")
	}

	adjusted := clean[4:] + clean[:4]
	ns := ""
	for _, char := range adjusted {
		// A-Z
		if int(char) >= 65 && int(char) <= 90 {
			ns = ns + strconv.Itoa(int(char)-55)
			continue
		}
		// 0-9
		if int(char) >= 48 && int(char) <= 57 {
			ns = ns + string(char)
			continue
		}
		// Anything else means this is not a well formed IBAN
		return "", errors.New("Badly formed IBAN")
	}

	// We now compute mod(97) of the above transformed numeric string
	iban := new(big.Int)
	if _, ok := iban.SetString(ns, 10); !ok {
		return "", errors.New("Badly formed IBAN")
	}
	ninetyseven := new(big.Int)
	ninetyseven.SetInt64(97)
	modulus := iban.Mod(iban, ninetyseven)

	if modulus.Int64() != 1 {
		return clean, errors.New("Invalid modulus")
	}

	c.countryCode = clean[0:2]
	c.ibanCheckDigits = clean[2:4]
	c.bban = clean[4:]

	return clean, nil

}
