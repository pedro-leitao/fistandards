// Package iban is a set of methods to validate and process IBAN representations.
// Original regular expressions taken from http://ht5ifv.serprest.pt/extensions/tools/IBAN/
package iban

import (
	"regexp"
	"strings"
)

// Checker is the representation of an IBAN validator
type Checker struct {
	regex             map[string]string
	compiledRegex     map[string]*regexp.Regexp
	compiledRegexNoCC map[string]*regexp.Regexp
}

// Initialize creates a new IBAN validator
func (c *Checker) Initialize() error {
	c.regex = make(map[string]string)
	c.compiledRegex = make(map[string]*regexp.Regexp)
	c.compiledRegexNoCC = make(map[string]*regexp.Regexp)

	c.regex["AL"] = `^AL\d{10}[0-9A-Z]{16}$`
	c.regex["AD"] = `^AD\d{10}[0-9A-Z]{12}$`
	c.regex["AT"] = `^AT\d{18}$`
	c.regex["BH"] = `^BH\d{2}[A-Z]{4}[0-9A-Z]{14}`
	c.regex["BE"] = `^BE\d{14}$`
	c.regex["BA"] = `^BA\d{18}$`
	c.regex["BG"] = `^BG\d{2}[A-Z]{4}\d{6}[0-9A-Z]{8}$`
	c.regex["HR"] = `^HR\d{19}$`
	c.regex["CY"] = `^CY\d{10}[0-9A-Z]{16}$`
	c.regex["CZ"] = `^CZ\d{22}$`
	c.regex["DK"] = `^DK\d{16}$|^FO\d{16}$|^GL\d{16}$`
	c.regex["DO"] = `^DO\d{2}[0-9A-Z]{4}\d{20}$`
	c.regex["EE"] = `^EE\d{18}$`
	c.regex["FI"] = `^FI\d{16}$`
	c.regex["FR"] = `^FR\d{12}[0-9A-Z]{11}\d{2}$`
	c.regex["GE"] = `^GE\d{2}[A-Z]{2}\d{16}$`
	c.regex["DE"] = `^DE\d{20}$`
	c.regex["GI"] = `^GI\d{2}[A-Z]{4}[0-9A-Z]{15}$`
	c.regex["GR"] = `^GR\d{9}[0-9A-Z]{16}$`
	c.regex["HU"] = `^HU\d{26}$`
	c.regex["IS"] = `^IS\d{24}$`
	c.regex["IE"] = `^IE\d{2}[A-Z]{4}\d{14}$`
	c.regex["IL"] = `^IL\d{21}$`
	c.regex["IT"] = `^IT\d{2}[A-Z]\d{10}[0-9A-Z]{12}$`
	//c.regex["KZ"] = `^[A-Z]{2}\d{5}[0-9A-Z]{13}$`
	c.regex["KW"] = `^KW\d{2}[A-Z]{4}22!$`
	c.regex["LV"] = `^LV\d{2}[A-Z]{4}[0-9A-Z]{13}$`
	c.regex["LB"] = `^LB\d{6}[0-9A-Z]{20}$`
	c.regex["LI"] = `^LI\d{7}[0-9A-Z]{12}$`
	c.regex["LT"] = `^LT\d{18}$`
	c.regex["LU"] = `^LU\d{5}[0-9A-Z]{13}$`
	c.regex["MK"] = `^MK\d{5}[0-9A-Z]{10}\d{2}$`
	c.regex["MT"] = `^MT\d{2}[A-Z]{4}\d{5}[0-9A-Z]{18}$`
	c.regex["MR"] = `^MR13\d{23}$`
	c.regex["MU"] = `^MU\d{2}[A-Z]{4}\d{19}[A-Z]{3}$`
	c.regex["MC"] = `^MC\d{12}[0-9A-Z]{11}\d{2}$`
	c.regex["ME"] = `^ME\d{20}$`
	c.regex["NL"] = `^NL\d{2}[A-Z]{4}\d{10}$`
	c.regex["NO"] = `^NO\d{13}$`
	c.regex["PL"] = `^PL\d{10}[0-9A-Z]{,16}n$`
	c.regex["PT"] = `^PT\d{23}$`
	c.regex["RO"] = `^RO\d{2}[A-Z]{4}[0-9A-Z]{16}$`
	c.regex["SM"] = `^SM\d{2}[A-Z]\d{10}[0-9A-Z]{12}$`
	c.regex["SA"] = `^SA\d{4}[0-9A-Z]{18}$`
	c.regex["RS"] = `^RS\d{20}$`
	c.regex["SK"] = `^SK\d{22}$`
	c.regex["SI"] = `^SI\d{17}$`
	c.regex["ES"] = `^ES\d{22}$`
	c.regex["SE"] = `^SE\d{22}$`
	c.regex["CH"] = `^CH\d{7}[0-9A-Z]{12}$`
	c.regex["TN"] = `^TN59\d{20}$`
	c.regex["TR"] = `^TR\d{7}[0-9A-Z]{17}$`
	c.regex["AE"] = `^AE\d{21}$`
	c.regex["GB"] = `^GB\d{2}[A-Z]{4}\d{14}$`

	for k, v := range c.regex {
		var err error
		if c.compiledRegex[k], err = regexp.Compile(v); err != nil {
			return err
		}
		if c.compiledRegexNoCC[k], err = regexp.Compile(`^` + v[3:]); err != nil {
			return err
		}
	}

	return nil
}

// Validate checks a string and returns a country code if it is a matching IBAN.
// If 'cleanup' is set to true the input string is cleaned for common patterns before checking, and is returned.
func (c *Checker) Validate(s string, cleanup bool) (string, string, bool) {
	if len(s) < 2 {
		return "", s, false
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
	if regex, ok := c.compiledRegex[s[0:2]]; ok {
		if regex.MatchString(s) {
			return s[0:2], s, true
		}
	}

	return "", s, false
}

// Guess checks a string (which doesn't include the country code) and returns a list of possible country codes if it is a matching IBAN.
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
	for k, v := range c.compiledRegexNoCC {
		if v.MatchString(s) {
			results[k] = k + s
		}
	}
	if len(results) > 0 {
		return results, true
	}

	return nil, false
}
