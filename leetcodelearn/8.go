package leetcodelearn

import "math"

func myAtoi(s string) int {

	isBegin := false
	isMinus := false
	r := 0
	if len(s) == 0 {
		return 0
	}

	for _, v := range s {
		if v >= 48 && v <= 57 {
			isBegin = true
			if r > math.MaxInt32/10 || (r == math.MaxInt32/10 && int(v-48) > math.MaxInt32%10) {
				r = math.MaxInt32
				if isMinus {
					r = math.MinInt32
				}
				return r
			}
			r = r*10 + int(v-48)
		} else if v == '-' {
			if isBegin {
				break
			} else {
				isMinus = true
				isBegin = true
			}
		} else if v == '+' {
			if isBegin {
				break
			}
			isBegin = true
		} else if v == ' ' && !isBegin {
			continue
		} else {
			break
		}
	}
	if isMinus {
		r *= -1
		if r < math.MinInt32 {
			r = math.MinInt32
		}
	}
	return r
}
