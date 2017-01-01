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
