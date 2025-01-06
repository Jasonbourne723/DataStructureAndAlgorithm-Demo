package leetcodelearn

import "math"

func reverse(x int) int {

	r := 0
	for x != 0 {
		m := x % 10
		x /= 10

		if r > math.MaxInt32/10 || (r == math.MaxInt32/10 && m > math.MaxInt32%10) {
			return 0
		}

		if r < math.MinInt32/10 || (r == math.MinInt32/10 && m < math.MinInt32%10) {
			return 0
		}

		r = r*10 + m

	}
	return r
}
