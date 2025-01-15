package leetcodelearn

func longestPalindrome(s string) string {
	begin := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if IsPalindrome(s[j : i+1]) {
				if i+1-j > maxLen {
					begin = j
					maxLen = i + 1 - j
				}
			}
		}
	}

	return s[begin : begin+maxLen]

}

func IsPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
