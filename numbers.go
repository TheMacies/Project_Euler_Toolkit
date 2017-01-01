package kit

import "math"

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

//GetFactors returns counted factors of a number
func GetFactors(number int) []int {
	factors := make([]int, number+1)
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			factors[i]++
			number /= i
			i = 1
		}
	}
	return factors
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

//ConcatNumbers specified in args into one big number
func ConcatNumbers(nums ...int) int {
	number := 0
	for i := range nums {
		digs := CountDigits(nums[i])
		number *= int(math.Pow10(digs))
		number += nums[i]
	}
	return number
}

//GeneratePrimesLookup generates a  n-length table such that i-th elem is true if i is prime number
func GeneratePrimesLookup(n int) []bool {
	primesLookup := make([]bool, n)
	numberCheck := make([]bool, n)

	for i := 2; i < n; i++ {
		if !numberCheck[i] {
			primesLookup[i] = true
			for j := i; j < n; j += i {
				numberCheck[j] = true
			}
		}
	}
	return primesLookup
}

func ReverseNumber(num int) int {
	a := 0
	for num != 0 {
		a *= 10
		a += num % 10
		num /= 10
	}
	return a
}

//NumberToArray converts number to an array
func NumberToArray(num int) []int {
	var comb []int
	if num == 0 {
		comb = append(comb, 0)
		return comb
	}
	for num > 0 {
		comb = append(comb, num%10)
		num = num / 10
	}
	return comb
}

func ArrayToNumber(comb []int) int {
	sum := 0
	multiplier := 1
	for i := range comb {
		sum += multiplier * comb[i]
		multiplier *= 10
	}
	return sum
}

//GeneratePermutations generates all permutations of given array
func GeneratePermutations(elems []int) [][]int {
	perms := make([][]int, 1)
	perms[0] = make([]int, len(elems))
	neutralNumber := 0
	for _, val := range elems {
		if neutralNumber >= val {
			neutralNumber = val - 1
		}
	}
	for i := range perms[0] {
		perms[0][i] = neutralNumber
	}
	for i := range elems {
		tempCombs := perms
		perms = make([][]int, 0)
		for j := range tempCombs {
			for k := 0; k < len(elems); k++ {
				tempComb := make([]int, len(elems))
				copy(tempComb, tempCombs[j])
				if tempComb[k] == neutralNumber {
					tempComb[k] = elems[i]
					perms = append(perms, tempComb)
				}
			}
		}
	}
	return perms
}

func GeneratePowerSet(numbers []int) [][]int {
	var res [][]int
	if len(numbers) == 1 {
		a := make([]int, 1)
		var b []int
		a[0] = numbers[0]
		res = append(res, a)
		res = append(res, b)
		return res
	}
	number := numbers[0]
	s1 := GeneratePowerSet(numbers[1:])
	s2 := GeneratePowerSet(numbers[1:])
	for i := range s1 {
		s1[i] = append(s1[i], number)
	}
	return append(s1, s2...)
}
