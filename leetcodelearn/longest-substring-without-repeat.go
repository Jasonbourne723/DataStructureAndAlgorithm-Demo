package leetcodelearn

func LengthOfLongestSubstring(s string) int {

	s1, s2 := "", ""

	for _, v := range s {

		index := -1
		for i, v1 := range s2 {
			if v == v1 {
				index = i
				break
			}
		}

		if index == -1 {
			runes2 := []rune(s2)
			runes2 = append(runes2, v)
			s2 = string(runes2)
		} else if index < len(s2)-1 {
			runes2 := []rune(s2[index+1:])
			runes2 = append(runes2, v)
			s2 = string(runes2)
		} else {
			s2 = string([]rune{v})
		}

		if len(s2) > len(s1) {
			s1 = s2
		}

	}

	return len(s1)
}
