package kit

import "testing"

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome("abbc") {
		t.Fail()
	}

	if !IsPalindrome("a") {
		t.Fail()
	}

	if !IsPalindrome("abgffgba") {
		t.Fail()
	}
}

func TestReverseString(t *testing.T) {
	if ReverseString("str") != "rts" {
		t.Fail()
	}
	if ReverseString("") != "" {
		t.Fail()
	}
}
