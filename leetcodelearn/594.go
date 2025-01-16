package leetcodelearn

import "sort"

func findLHS(nums []int) int {

	sort.Ints(nums)

	begin, ans := 0, 0

	for end, v := range nums {
		if v > nums[begin]+1 {
			begin++
		}
		if v-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return ans
}
