package bban

import (
	"testing"
)

func TestValidate(t *testing.T) {
	var c Checker

	got := c.Initialize()
	if got != nil {
		t.Errorf("Initialize() = %v; want nil", got.Error())
	}

	gotclean, gotmatched := c.Validate("3456789012345678", "AT", false)
	wantclean := "3456789012345678"
	wantmatched := true
	if gotmatched != wantmatched || gotclean != wantclean {
		t.Errorf("Validate(3456789012345678) = %v,%v; want %v,%v", gotclean, gotmatched, wantclean, wantmatched)
	}

	gotclean, gotmatched = c.Validate("  3-456.789-0123.45678		", "AT", true)
	wantclean = "3456789012345678"
	wantmatched = true
	if gotmatched != wantmatched || gotclean != wantclean {
		t.Errorf("Validate(  3-456.789-0123.45678		) = %v,%v; want %v,%v", gotclean, gotmatched, wantclean, wantmatched)
	}

	gotclean, gotmatched = c.Validate("ZZ123", "AT", false)
	wantclean = "ZZ123"
	wantmatched = false
	if gotmatched != wantmatched || gotclean != wantclean {
		t.Errorf("Validate(ZZ123) = %v,%v; want %v,%v", gotclean, gotmatched, wantclean, wantmatched)
	}

	var gotcclist map[string]string

	gotcclist, gotmatched = c.Guess("  3-456.789-0123.45678		", true)
	wantmatched = true
	if gotmatched != wantmatched {
		t.Errorf("Guess(  3-456.789-0123.45678		) = %v,%v; want %v", gotcclist, gotmatched, wantmatched)
	}

	gotcclist, gotmatched = c.Guess("SRLG60837158918739", true)
	wantmatched = true
	if _, ok := gotcclist["GB"]; !ok {
		t.Errorf("Guess(SRLG60837158918739) = %v,%v; want %v", gotcclist, gotmatched, wantmatched)
	}

}
