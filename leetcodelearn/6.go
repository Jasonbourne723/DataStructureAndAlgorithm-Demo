package leetcodelearn

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	ans := make([][]byte, numRows)

	prev := 0
	j := 0
	for i := 0; i < len(s); i++ {
		ans[j] = append(ans[j], s[i])
		temp := j
		if j == 0 {
			j++
		} else if j == numRows-1 {
			j--
		} else {
			if j > prev {
				j++
			} else {
				j--
			}
		}
		prev = temp
	}

	s1 := make([]byte, 0, len(s))

	for _, v := range ans {
		s1 = append(s1, v...)
	}
	return string(s1)
}
