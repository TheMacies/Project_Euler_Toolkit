package kit

// CountDigits  of an integer
func CountDigits(number int) int {
	if number == 0 {
		return 1
	}
	if number < 0 {
		number = -number
	}

	digits := 0
	for number > 0 {
		digits++
		number /= 10
	}
	return digits
}

//NPandigital checks if given number is n-pandigital
func NPandigital(number, n int) bool {
	digits := CountDigits(number)
	if digits != n {
		return false
	}
	count := make([]int, n+1)
	for number > 0 {
		count[number%10]++
		number /= 10
	}

	if count[0] != 0 {
		return false
	}

	for i := 1; i < n; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return true
}
