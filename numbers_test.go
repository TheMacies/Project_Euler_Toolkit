package kit

import (
	"testing"
)

func TestCountDigits(t *testing.T) {
	if CountDigits(0) != 1 {
		t.Fail()
	}
	if CountDigits(-50) != CountDigits(12) {
		t.Fail()
	}
	if CountDigits(1203) != 4 {
		t.Fail()
	}
}

func TestNPandigital(t *testing.T) {
	if !NPandigital(1423, 4) {
		t.Fail()
	}

	if NPandigital(421256, 6) {
		t.Fail()
	}

	if NPandigital(012345, 5) {
		t.Fail()
	}

	if NPandigital(20345, 5) {
		t.Fail()
	}

	if !NPandigital(324196578, 9) {
		t.Fail()
	}
}

func TestConcatNumbers(t *testing.T) {
	if ConcatNumbers(2, 3, 4, 5) != 2345 {
		t.Fail()
	}
	if ConcatNumbers(0, 2, 0, 2) != 202 {
		t.Fail()
	}
}
