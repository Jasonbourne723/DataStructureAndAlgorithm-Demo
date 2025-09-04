package leetcodelearn

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	var r []string

	var t uint16

	var c int

	var backtrack func(int)
	backtrack = func(p int) {

		if c == turnedOn {
			s := getstr(t)
			r = append(r, s)
			return
		}

		for i := p; i <= 10; i++ {
			t |= 1 << i
			c++
			backtrack(i + 1)
			t = t &^ (1 << i)
			c--
		}
	}

	backtrack(0)
	return r
}

func getstr(p uint16) string {

	h := p >> 6
	m := (p << 10) >> 10
	s := fmt.Sprintf("%d:%02d", h, m)
	return s
}
