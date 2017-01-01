package kit

import "testing"

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

func TestGeneratePrimesLookup(t *testing.T) {
	a := GeneratePrimesLookup(1000)
	if a[1] == true {
		t.Fail()
	}

	if a[6] == true {
		t.Fail()
	}

	if a[121] == true {
		t.Fail()
	}

	if a[2] == false {
		t.Fail()
	}

	if a[967] == false {
		t.Fail()
	}
}

func TestReverseNumber(t *testing.T) {
	if ReverseNumber(2) != 2 {
		t.Fail()
	}

	if ReverseNumber(210) != 12 {
		t.Fail()
	}

	if ReverseNumber(1234556) != 6554321 {
		t.Fail()
	}
}

func TestNumberToArray(t *testing.T) {
	a := NumberToArray(120)
	if a[0] != 0 {
		t.Fail()
	}
	if a[1] != 2 {
		t.Fail()
	}
	if a[2] != 1 {
		t.Fail()
	}

	a = NumberToArray(0)
	if a[0] != 0 {
		t.Fail()
	}
}

func TestArrayToNumber(t *testing.T) {
	var a []int
	a = make([]int, 3)
	a[0] = 9
	a[1] = 2
	a[2] = 1
	if ArrayToNumber(a) != 129 {
		t.Fail()
	}
	if ArrayToNumber(NumberToArray(10)) != 10 {
		t.Fail()
	}
}

func TestGeneratePermutations(t *testing.T) {
	a := []int{-1, 2}
	p := GeneratePermutations(a)
	el1 := p[0]
	el2 := p[1]
	if el1[0] != -1 || el1[1] != 2 {
		t.Fail()
	}
	if el2[0] != 2 || el2[1] != -1 {
		t.Fail()
	}
}

func TestGetFactors(t *testing.T) {
	f := GetFactors(2)
	if f[2] != 1 {
		t.Fail()
	}
	if f[1] != 0 {
		t.Fail()
	}
	f = GetFactors(121)
	if f[11] != 2 {
		t.Fail()
	}
	f = GetFactors(1)
	if f[1] != 0 {
		t.Fail()
	}
}
