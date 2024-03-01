package leetcodelearn

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Ints(s)
	sort.Ints(g)

	num := 0

	ls := len(s)
	j := 0
	for _, gv := range g {
		for j < ls {
			if s[j] >= gv {
				j++
				num++
				break
			}
			j++
		}
	}

	return num

}
