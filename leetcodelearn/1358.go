package leetcodelearn

func NumberOfSubstrings(s string) int {
	var c int
	m := make(map[byte]int)
	right := 0
	for i := 0; i <= len(s)-3; i++ {
		if i > 0 {
			m[s[i]]--
		}
		isExist := false
		for ; right < len(s); right++ {
			m[s[right]]++
			if m['a'] > 0 && m['b'] > 0 && m['c'] > 0 {
				c += len(s) - right
				isExist = true
				break
			}
		}
		if !isExist {
			break
		}
		if right < len(s) {
			m[s[right]]--
		}
	}
	return c
}
