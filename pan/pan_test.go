package pan

import (
	"testing"
)

func TestGetIssuer(t *testing.T) {
	var c Pan

	_, err := c.Validate("5460976249685093")
	wantissuer := "Mastercard"
	if c.Issuer != wantissuer || err != nil {
		t.Errorf("GetIssuer(5460976249685093) = %v, %v; want %v", c.Issuer, err, wantissuer)
	}

	_, err = c.Validate("5574351064815121")
	wantissuer = "Mastercard"
	if c.Issuer != wantissuer || err != nil {
		t.Errorf("GetIssuer(5574351064815121) = %v, %v; want %v", c.Issuer, err, wantissuer)
	}

	_, err = c.Validate("5574351064815128")
	wantissuer = "Mastercard"
	if c.Issuer != wantissuer || err == nil {
		t.Errorf("GetIssuer(5574351064815128) = %v, %v; want %v", c.Issuer, err, wantissuer)
	}
}
