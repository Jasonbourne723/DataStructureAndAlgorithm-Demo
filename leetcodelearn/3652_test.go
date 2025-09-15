package leetcodelearn

import "testing"

func TestMaxProfit(t *testing.T) {
	s := MaxProfit([]int{4, 7, 13}, []int{-1, -1, 0}, 2)
	print(s)
}
