package iban

import (
	"testing"
)

func TestVerify(t *testing.T) {
	var c Iban

	gotclean, goterror := c.Validate("GB82 WEST 1234 5698 7654 32")
	wantclean := "GB82WEST12345698765432"
	if gotclean != wantclean || goterror != nil {
		t.Errorf("Verify(GB82 WEST 1234 5698 7654 32) = %v, %v; want %v, %v", gotclean, goterror, wantclean, nil)
	}

	gotclean, goterror = c.Validate("GB82 WEST 1234 6698 7654 32")
	wantclean = "GB82WEST12346698765432"
	if gotclean != wantclean || goterror == nil {
		t.Errorf("Verify(GB82 WEST 1234 6698 7654 32) = %v, %v; want %v, %v", gotclean, goterror.Error(), wantclean, "Invalid modulus")
	}
}
