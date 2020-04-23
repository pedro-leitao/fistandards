package iban

import (
	"testing"
)

func TestValidate(t *testing.T) {
	var c Checker

	got := c.Initialize()
	if got != nil {
		t.Errorf("Initialize() = %v; want nil", got.Error())
	}

	gotcc, gotclean, gotmatched := c.Validate("AT123456789012345678", false)
	wantcc := "AT"
	wantclean := "AT123456789012345678"
	wantmatched := true
	if gotcc != wantcc || gotmatched != wantmatched || gotclean != wantclean {
		t.Errorf("Validate(AT123456789012345678) = %v,%v,%v; want %v,%v,%v", gotcc, gotclean, gotmatched, wantcc, wantclean, wantmatched)
	}

	gotcc, gotclean, gotmatched = c.Validate("  A-T 123-456.789-0123.45678		", true)
	wantcc = "AT"
	wantclean = "AT123456789012345678"
	wantmatched = true
	if gotcc != wantcc || gotmatched != wantmatched || gotclean != wantclean {
		t.Errorf("Validate(  A-T 123-456.789-0123.45678		) = %v,%v,%v; want %v,%v,%v", gotcc, gotclean, gotmatched, wantcc, wantclean, wantmatched)
	}

	gotcc, gotclean, gotmatched = c.Validate("ZZ123", false)
	wantcc = ""
	wantclean = "ZZ123"
	wantmatched = false
	if gotcc != wantcc || gotmatched != wantmatched || gotclean != wantclean {
		t.Errorf("Validate(ZZ123) = %v,%v,%v; want %v,%v,%v", gotcc, gotclean, gotmatched, wantcc, wantclean, wantmatched)
	}

	var gotcclist map[string]string

	gotcclist, gotmatched = c.Guess("  123-456.789-0123.45678		", true)
	wantmatched = true
	if gotmatched != wantmatched {
		t.Errorf("Guess(  123-456.789-0123.45678		) = %v,%v; want %v", gotcclist, gotmatched, wantmatched)
	}

	gotcclist, gotmatched = c.Guess("62SRLG60837158918739", true)
	wantmatched = true
	if _, ok := gotcclist["GB"]; !ok {
		t.Errorf("Guess(62SRLG60837158918739) = %v,%v; want %v", gotcclist, gotmatched, wantmatched)
	}

}
