package leetcodelearn

func Partition(s string) [][]string {
	r := [][]string{}
	var p []string

	var backtrack func(int)

	backtrack = func(index int) {
		if index == len(s) {
			t := make([]string, len(p))
			copy(t, p)
			r = append(r, t)
			return
		}

		for i := index; i < len(s); i++ {
			if isPalindrome(s[index : i+1]) {
				p = append(p, s[index:i+1])
				backtrack(i + 1)
				p = p[:len(p)-1]
			}
		}
	}
	backtrack(0)
	return r
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
