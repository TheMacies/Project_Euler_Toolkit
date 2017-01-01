package kit

func IsPalindrome(st string) bool {
	for i := 0; i < len(st)/2; i++ {
		if st[i] != st[len(st)-1-i] {
			return false
		}
	}
	return true
}

func ReverseString(st string) string {
	var a string
	for i := range st {
		a += string(st[len(st)-1-i])
	}
	return a
}
