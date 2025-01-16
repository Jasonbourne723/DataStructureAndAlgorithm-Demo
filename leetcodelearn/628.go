package leetcodelearn

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt, math.MaxInt
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt

	for _, v := range nums {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}

		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	sort.Ints(nums)

	m := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {

		if _, exist := m[nums[i]]; exist {
			ans[0] = nums[i]
		}
		m[nums[i]] = true

		if nums[i] != i+1 && ans[1] == 0 {
			ans[1] = i + 1
		}
		if ans[0] != 0 && ans[1] != 0 {
			break
		}
	}
	return ans
}

func removeDuplicates(nums []int) int {

	if len(nums) < 3 {
		return len(nums)
	}

	s := 2
	for f := 2; f < len(nums); f++ {
		if nums[f] != nums[f-2] {
			nums[s] = nums[f]
			f--
			s++
		}
	}
	return s
}
