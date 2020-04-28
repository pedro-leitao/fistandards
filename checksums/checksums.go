package checksums

import (
	"errors"
	"math/big"
	"strconv"
)

// Mod97 takes a string which represents a long integer, and checks the modulus of s/97 is 1
func Mod97(s string) error {
	// We now compute mod(97) of the above numeric string
	n := new(big.Int)
	if _, ok := n.SetString(s, 10); !ok {
		return errors.New("Cannot compute modulus")
	}
	ninetyseven := new(big.Int)
	ninetyseven.SetInt64(97)
	modulus := n.Mod(n, ninetyseven)

	if modulus.Int64() != 1 {
		return errors.New("Invalid modulus")
	}

	return nil
}

// Mod10 is an implementation of the Luhn algorithm, which takes a string representing a long integer,
// and checks the modulus of s/10, using the last digit as the check digit.
func Mod10(s string) error {
	runes := []rune(s)
	sum := 0
	isSecond := false
	for k := len(runes) - 1; k >= 0; k-- {
		if isSecond {
			switch runes[k] {
			case '1':
				runes[k] = '2'
			case '2':
				runes[k] = '4'
			case '3':
				runes[k] = '6'
			case '4':
				runes[k] = '8'
			case '5':
				runes[k] = '1'
			case '6':
				runes[k] = '3'
			case '7':
				runes[k] = '5'
			case '8':
				runes[k] = '7'
			case '9':
				runes[k] = '9'
			}
		}
		digit, ok := strconv.Atoi(string(runes[k]))
		if ok != nil {
			return errors.New("Cannot compute modulus")
		}
		sum = sum + digit
		isSecond = !isSecond
	}

	if sum%10 == 0 {
		return nil
	}

	return errors.New("Invalid modulus")
}
