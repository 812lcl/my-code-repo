package leetcode

import "testing"

func TestNumDecodings(t *testing.T) {
	// Test case 1: Empty string
	if res := numDecodings(""); res != 0 {
		t.Errorf("numDecodings() = %d, want 0", res)
	}

	// Test case 2: Single digit
	if res := numDecodings("1"); res != 1 {
		t.Errorf("numDecodings() = %d, want 1", res)
	}

	// Test case 3: Two digits
	if res := numDecodings("12"); res != 2 {
		t.Errorf("numDecodings() = %d, want 2", res)
	}

	// Test case 4: Three digits
	if res := numDecodings("226"); res != 3 {
		t.Errorf("numDecodings() = %d, want 3", res)
	}

	// Test case 5: Invalid digits
	if res := numDecodings("0"); res != 0 {
		t.Errorf("numDecodings() = %d, want 0", res)
	}

	// Test case 6: Leading zeros
	if res := numDecodings("01"); res != 0 {
		t.Errorf("numDecodings() = %d, want 0", res)
	}
}
