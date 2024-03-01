package leetcodelearn

import "sort"

func findContentChildren1(g []int, s []int) int {

	sort.Ints(s)
	sort.Ints(g)
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		g[i], g[j] = g[j], g[i]
	}

	num := 0

	for _, sv := range s {
		for j, gv := range g {

			if sv >= gv {
				num++

				if j >= len(g) {
					g = g[:j]
				} else {
					g = append(g[:j], g[j+1:]...)
				}
				break

			}

		}
	}

	return num

}

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
