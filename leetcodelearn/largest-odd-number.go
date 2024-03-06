package leetcodelearn

func largestOddNumber(num string) string {

	l := len(num)

	index := -1

	for i := l - 1; i >= 0; i-- {
		b := num[i] - '0'
		if b%2 == 1 {
			index = i
			break
		}

	}
	if index == -1 {
		return ""
	} else {
		return num[:index+1]
	}

}
