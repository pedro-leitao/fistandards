package pan

import (
	"testing"
)

func TestGetIssuer(t *testing.T) {
	var c Pan

	_, _ = c.Set("5460976249685093")
	wantissuer := "Mastercard"
	if c.Issuer != wantissuer {
		t.Errorf("GetIssuer(5460976249685093) = %v; want %v", c.Issuer, wantissuer)
	}

	_, _ = c.Set("5574351064815121")
	wantissuer = "Mastercard"
	if c.Issuer != wantissuer {
		t.Errorf("GetIssuer(5574351064815121) = %v; want %v", c.Issuer, wantissuer)
	}
}
