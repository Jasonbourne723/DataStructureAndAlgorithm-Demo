package leetcodelearn

import (
	"strconv"
)

func maxDiff1(num int) int {

	var d = 1
	for {
		r := num / d
		if r >= 10 {
			d *= 10

		} else {
			break
		}
	}

	a := 0
	b := 0

	d1 := d
	num1 := num
	x1 := -1
	for {
		r := num1 / d1
		if r == 9 {
			a += r * d1
		} else {
			if x1 == -1 {
				x1 = r
			}
			if r == x1 {
				a += 9 * d1
			} else {
				a += r * d1
			}
		}
		if d1 == 1 {
			break
		}

		num1 %= d1
		d1 /= 10
	}

	d2 := d
	num2 := num
	i := 0
	x2 := -1
	y := 0

	first := 1
	for {
		r := num2 / d2

		if i == 0 {
			if r == 1 {
				b += r * d2
				first = 1
			} else {
				b += 1 * d2
				if x2 == -1 {
					x2 = r
					y = 1
				}
				first = r
			}
		} else {
			if r != 0 {
				if r == 1 && first == 1 {
					b += r * d2
				} else {
					if x2 == -1 {
						x2 = r
						y = 0
					}
					if r != x2 {
						b += r * d2
					} else {
						b += y * d2
					}
				}
			}
		}
		if d2 == 1 {
			break
		}

		i++
		num2 %= d2
		d2 /= 10

		if num2 == 0 {
			break
		}
	}

	return a - b
}

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

	for i, _ := range bytes {
		if x == int(bytes[i]-'0') {
			bytes[i] = byte('0' + y)
		}
	}
	return string(bytes)
}
