// Package bban is a set of methods to validate and process BBAN representations.
// Original regular expressions taken from http://ht5ifv.serprest.pt/extensions/tools/IBAN/
package bban

import (
	"regexp"
	"strings"
)

// Checker is the representation of a BBAN validator
type Checker struct {
	regex             map[string]string
	compiledRegex     map[string]*regexp.Regexp
}

// Initialize creates a new IBAN validator
func (c *Checker) Initialize() error {
	c.regex = make(map[string]string)
	c.compiledRegex = make(map[string]*regexp.Regexp)

	c.regex["AL"] = `^[0-9A-Z]{16}$`
	c.regex["AD"] = `^[0-9A-Z]{12}$`
	c.regex["AT"] = `^\d{16}$`
	c.regex["BH"] = `^[A-Z]{4}[0-9A-Z]{14}`
	c.regex["BE"] = `^\d{12}$`
	c.regex["BA"] = `^\d{16}$`
	c.regex["BG"] = `^[A-Z]{4}\d{6}[0-9A-Z]{8}$`
	c.regex["HR"] = `^\d{17}$`
	c.regex["CY"] = `^\d{8}[0-9A-Z]{16}$`
	c.regex["CZ"] = `^\d{20}$`
	c.regex["DK"] = `^\d{14}$`
	c.regex["DO"] = `^[0-9A-Z]{4}\d{20}$`
	c.regex["EE"] = `^\d{16}$`
	c.regex["FI"] = `^\d{14}$`
	c.regex["FR"] = `^\d{10}[0-9A-Z]{11}\d{2}$`
	c.regex["GE"] = `^[A-Z]{2}\d{16}$`
	c.regex["DE"] = `^\d{18}$`
	c.regex["GI"] = `^[A-Z]{4}[0-9A-Z]{15}$`
	c.regex["GR"] = `^\d{7}[0-9A-Z]{16}$`
	c.regex["HU"] = `^\d{24}$`
	c.regex["IS"] = `^\d{22}$`
	c.regex["IE"] = `^[A-Z]{4}\d{14}$`
	c.regex["IL"] = `^\d{19}$`
	c.regex["IT"] = `^[A-Z]\d{10}[0-9A-Z]{12}$`
	//c.regex["KZ"] = `^[A-Z]{2}\d{5}[0-9A-Z]{13}$`
	c.regex["KW"] = `^[A-Z]{4}[0-9A-Z]{22}$`
	c.regex["LV"] = `^[A-Z]{4}[0-9A-Z]{13}$`
	c.regex["LB"] = `^\d{4}[0-9A-Z]{20}$`
	c.regex["LI"] = `^\d{4}[0-9A-Z]{12}$`
	c.regex["LT"] = `^\d{16}$`
	c.regex["LU"] = `^\d{3}[0-9A-Z]{13}$`
	c.regex["MK"] = `^\d{3}[0-9A-Z]{10}\d{2}$`
	c.regex["MT"] = `^[A-Z]{4}\d{5}[0-9A-Z]{18}$`
	c.regex["MR"] = `^\d{23}$`
	c.regex["MU"] = `^[A-Z]{4}\d{19}[A-Z]{3}$`
	c.regex["MC"] = `^\d{10}[0-9A-Z]{11}\d{2}$`
	c.regex["ME"] = `^\d{18}$`
	c.regex["NL"] = `^[A-Z]{4}\d{10}$`
	c.regex["NO"] = `^\d{11}$`
	c.regex["PL"] = `^\d{24}$`
	c.regex["PT"] = `^\d{21}$`
	c.regex["RO"] = `^[A-Z]{4}[0-9A-Z]{16}$`
	c.regex["SM"] = `^[A-Z]\d{10}[0-9A-Z]{12}$`
	c.regex["SA"] = `^\d{2}[0-9A-Z]{18}$`
	c.regex["RS"] = `^\d{18}$`
	c.regex["SK"] = `^\d{20}$`
	c.regex["SI"] = `^\d{15}$`
	c.regex["ES"] = `^\d{20}$`
	c.regex["SE"] = `^\d{20}$`
	c.regex["CH"] = `^\d{5}[0-9A-Z]{12}$`
	c.regex["TN"] = `^\d{20}$`
	c.regex["TR"] = `^\d{5}[0-9A-Z]{17}$`
	c.regex["AE"] = `^\d{19}$`
	c.regex["GB"] = `^[A-Z]{4}\d{14}$`

	for k, v := range c.regex {
		var err error
		if c.compiledRegex[k], err = regexp.Compile(v); err != nil {
			return err
		}
	}

	return nil
}

// Validate checks a string and country code and returns true if it is a matching IBAN.
// If 'cleanup' is set to true the input string is cleaned for common patterns before checking, and is returned.
func (c *Checker) Validate(s string, countrycode string, cleanup bool) (string, bool) {
	if len(s) < 2 {
		return s, false
	}

	if cleanup {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, "-", "")
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "\t", "")
		s = strings.ToUpper(s)
	}

	// Try to match by country code
	if regex, ok := c.compiledRegex[countrycode]; ok {
		if regex.MatchString(s) {
			return s, true
		}
	}

	return s, false
}

// Guess checks a string and returns a list of possible country codes if it is a matching IBAN.
// If 'cleanup' is set to true the input string is cleaned for common patterns before checking, and is returned.
func (c *Checker) Guess(s string, cleanup bool) (map[string]string, bool) {
	if len(s) < 2 {
		return nil, false
	}

	if cleanup {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, "-", "")
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, "\t", "")
		s = strings.ToUpper(s)
	}

	// Try to guess country code
	results := make(map[string]string)
	for k, v := range c.compiledRegex {
		if v.MatchString(s) {
			results[k] = s
		}
	}
	if len(results) > 0 {
		return results, true
	}

	return nil, false
}
