// Package pan is a set of methods to validate and process PAN representations.
// See https://en.wikipedia.org/wiki/Payment_card_number for structure details.
package pan

import (
	"errors"
	"strconv"

	"github.com/pedro-leitao/fistandards/checksums"
	"github.com/pedro-leitao/fistandards/utils"
)

// Mapping of the possible initial digits for various card issuers, should be read as:
// <from>-<to> <issuer>[<min length>-<max length>] <validation algorithm>
var issuersMapping = map[string]int{
	"1-1 UATP [15-15] luhn":                            1,
	"4-4 Visa [16-16] luhn":                            1,
	"34-34 American Express [15-15] luhn":              2,
	"37-37 American Express [15-15] luhn":              2,
	"31-31 China T-Union [19-19] luhn":                 2,
	"62-62 China UnionPay [16-19] luhn":                2,
	"81-81 China UnionPay [16-19] luhn":                2,
	"36-36 Diners Club International [14-19] luhn":     2,
	"38-39 Diners Club International [16-19] luhn":     2,
	"60-60 RuPay [16-16] luhn":                         2,
	"50-50 Maestro [12-19] luhn":                       2,
	"56-69 Maestro [12-19] luhn":                       2,
	"51-55 Mastercard [16-16] luhn":                    2,
	"64-64 Discover Card [16-19] luhn":                 2,
	"65-65 Discover Card [16-19] luhn":                 2,
	"636-636 Interpayment [16-19] luhn":                3,
	"300-305 Diners Club International [16-19] luhn":   3,
	"637-639 InstaPayment [16-16] luhn":                3,
	"5610-5610 Bankcard [16-16] luhn":                  4,
	"2014-2014 Diners Club enRoute [15-15] none":       4,
	"2149-2149 Diners Club enRoute [15-15] none":       4,
	"3095-3095 Diners Club International [16-19] luhn": 4,
	"6011-6011 Discover Card [16-19] luhn":             4,
	"6521-6522 RuPay [16-16] luhn":                     4,
	"3528–3589 JCB [16-19] luhn":                       4,
	"6304-6304 Laser [16-19] luhn":                     4,
	"6706-6706 Laser [16-19] luhn":                     4,
	"6771-6771 Laser [16-19] luhn":                     4,
	"6709-6709 Laser [16-19] luhn":                     4,
	"6759-6759 Maestro UK [12-19] luhn":                4,
	"5019-5019 Dankort [16-16] luhn":                   4,
	"4571-4571 Dankort [16-16] luhn":                   4,
	"2200-2204 MIR [16-16] luhn":                       4,
	"2221-2720 Mastercard [16-16] luhn":                4,
	"6334-6334 Solo [16-16] luhn":                      4,
	"6334-6334 Solo [18-19] luhn":                      4,
	"6767-6767 Solo [16-16] luhn":                      4,
	"6767-6767 Solo [18-19] luhn":                      4,
	"4903-4903 Switch [16-16] luhn":                    4,
	"4903-4903 Switch [18-19] luhn":                    4,
	"4905-4905 Switch [16-16] luhn":                    4,
	"4905-4905 Switch [18-19] luhn":                    4,
	"4911-4911 Switch [16-16] luhn":                    4,
	"4911-4911 Switch [18-19] luhn":                    4,
	"4936-4936 Switch [16-16] luhn":                    4,
	"4936-4936 Switch [18-19] luhn":                    4,
	"6333-6333 Switch [16-16] luhn":                    4,
	"6333-6333 Switch [18-19] luhn":                    4,
	"6759-6759 Switch [16-16] luhn":                    4,
	"6759-6759 Switch [18-19] luhn":                    4,
	"560221–560225 Bankard [16-16] luhn":               6,
	"622126-622925 Discover Card [16-19] luhn":         6,
	"624000-626999 Discover Card [16-19] luhn":         6,
	"628200-628899 Discover Card [16-19] luhn":         6,
	"676770-676770 Maestro UK [12-19] luhn":            6,
	"676774-676774 Maestro UK [12-19] luhn":            6,
	"564182-564182 Switch [16-16] luhn":                6,
	"564182-564182 Switch [18-19] luhn":                6,
	"633110-633110 Switch [16-16] luhn":                6,
	"633110-633110 Switch [18-19] luhn":                6,
	"979200–979289 Troy [16-16] luhn":                  6,
	"506099–506198 Verve [16-16] luhn":                 6,
	"506099–506198 Verve [19-19] luhn":                 6,
	"650002–650027 Verve [16-16] luhn":                 6,
	"650002–650027 Verve [19-19] luhn":                 6,
	"357111-357111 LankaPay [16-16] luhn":              6,
	"6054740-6054744 NPS Pridnestrovie [16-16] luhn":   7,
}

// Pan is the representation of a PAN value
type Pan struct {
	Pan       string
	Issuer    string
	Iin       int
	Algorithm string
	valid     bool
}

// getIssuer returns the card issuer for a given PAN
func (c *Pan) getIssuer(s string) (err error) {
	var re = `(?mU)^(?P<rangeFrom>[0-9]+)[^0-9](?P<rangeTo>[0-9]+)\s(?P<issuer>.*)\s\[(?P<lenFrom>[0-9]+)[^0-9](?P<lenTo>[0-9]+)\]\s(?P<algorithm>none|luhn)$`
	// Traverse all issuers, find the correct one for this string
	for k, v := range issuersMapping {
		// Parse the notation
		paramsMap := utils.GetParams(re, k)

		var rangeFrom, rangeTo, lenFrom, lenTo int
		var ok error
		var issuer, algorithm string

		if rangeFrom, ok = strconv.Atoi(paramsMap["rangeFrom"]); ok != nil {
			return errors.New("Failed to convert to Integer")
		}
		if rangeTo, ok = strconv.Atoi(paramsMap["rangeTo"]); ok != nil {
			return errors.New("Failed to convert to Integer")
		}
		if lenFrom, ok = strconv.Atoi(paramsMap["lenFrom"]); ok != nil {
			return errors.New("Failed to convert to Integer")
		}
		if lenTo, ok = strconv.Atoi(paramsMap["lenTo"]); ok != nil {
			return errors.New("Failed to convert to Integer")
		}
		issuer = paramsMap["issuer"]
		algorithm = paramsMap["algorithm"]

		iin, _ := strconv.Atoi(s[0:v])
		if iin >= rangeFrom && iin <= rangeTo && len(s) >= lenFrom && len(s) <= lenTo {
			c.Issuer = issuer
			c.Iin = iin
			c.Algorithm = algorithm
			return nil
		}
	}

	return errors.New("No issuer match found")
}

// Validate a given PAN string, returning a normalized
// representation and an error in case verification failed. Once validated, the calling object
// will have additional information regarding the PAN.
func (c *Pan) Validate(s string) (norm string, err error) {

	if len(s) < 8 {
		return s, errors.New("Invalid length")
	}

	clean := utils.Clean(s)
	if err := c.getIssuer(clean); err != nil {
		return clean, err
	}

	if c.Algorithm == "luhn" {
		if err := checksums.Mod10(clean); err != nil {
			c.valid = true
			return clean, err
		}
	}

	return clean, nil
}
