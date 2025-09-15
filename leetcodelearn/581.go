package leetcodelearn

import "math"

func findUnsortedSubarray(nums []int) int {
	var s [][]int
	begin := math.MaxInt
	var end int
	m := 0
	for i := 0; i < len(nums); i++ {
		if len(s) == 0 || nums[i] > s[len(s)-1][0] {
			s = append(s, []int{nums[i], i})
			if nums[i] > m {
				m = nums[i]
			}
		} else {
			for len(s) > 0 {
				if nums[i] < s[len(s)-1][0] {
					l := s[len(s)-1][1]
					if l < begin {
						begin = l
					}
					s = s[:len(s)-1]
					end = i
				} else if nums[i] < m {
					end = i
				} else {
					break
				}
			}

		}
	}
	if end == 0 {
		return 0
	}
	return end - begin + 1
}
