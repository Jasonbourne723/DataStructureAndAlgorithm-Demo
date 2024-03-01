package leetcodelearn

import (
	"strconv"
)

func maxDiff(num int) int {

	max := num
	min := num

	str := strconv.Itoa(num)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {

			r := Replace(str, x, y)
			if r[0] != '0' {
				rNum, _ := strconv.Atoi(r)

				if rNum > max {
					max = rNum
				}
				if rNum < min {
					min = rNum
				}
			}
		}
	}

	return max - min
}

func Replace(num string, x int, y int) string {

	bytes := []byte(num)

	for i := range bytes {
		if x == int(bytes[i]-'0') {
			bytes[i] = byte('0' + y)
		}
	}
	return string(bytes)
}
