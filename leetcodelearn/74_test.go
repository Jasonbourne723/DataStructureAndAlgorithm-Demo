package leetcodelearn

import "testing"

func TestSearchMatrix(t *testing.T) {
	a := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	r := SearchMatrix(a, 3)
	print(r)

}
