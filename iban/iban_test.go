package iban

import (
	"testing"
)

func TestVerify(t *testing.T) {
	var c Iban

	gotclean, gotverified := c.Set("GB82 WEST 1234 5698 7654 32")
	wantclean := "GB82WEST12345698765432"
	wantverified := true
	if gotclean != wantclean || gotverified != wantverified {
		t.Errorf("Verify(GB82 WEST 1234 5698 7654 32) = %v, %v; want %v, %v", gotclean, gotverified, wantclean, wantverified)
	}

	gotclean, gotverified = c.Set("GB82 WEST 1234 6698 7654 32")
	wantclean = "GB82WEST12346698765432"
	wantverified = false
	if gotclean != wantclean || gotverified != wantverified {
		t.Errorf("Verify(GB82 WEST 1234 6698 7654 32) = %v, %v; want %v, %v", gotclean, gotverified, wantclean, wantverified)
	}
}
