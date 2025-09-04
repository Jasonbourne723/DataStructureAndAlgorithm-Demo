package leetcodelearn

import "testing"

func TestSubsets(t *testing.T) {
	s := []int{1, 2, 3}
	r := Subsets(s)
	print(r)
}
