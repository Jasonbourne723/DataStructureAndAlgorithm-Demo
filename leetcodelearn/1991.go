package leetcodelearn

import "math/rand"

// 自己写的
func pivotIndex(nums []int) int {
	rightSum := 0

	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

//官方正解
func pivotIndex1(nums []int) int {

	total := 0
	for _, v := range nums {
		total += v
	}

	sum := 0
	for i, v := range nums {
		if sum*2+v == total {
			return i
		}
		sum += v
	}
	return -1

}

func searchInsert(nums []int, target int) int {
	return search(nums, 0, len(nums)-1, target)
}

func search(nums []int, left, right, target int) int {

	if nums[left] >= target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if nums[right] < target {
		return right + 1
	}
	m := (right-left)/2 + left
	if target == nums[m] {
		return m
	}

	if target > nums[m] {
		return search(nums, m+1, right, target)
	} else {
		return search(nums, left, m-1, target)
	}

}

func merge(intervals [][]int) [][]int {
	intervals = sort1(intervals)
	j := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[j][1] {
			j++
			intervals[j] = intervals[i]

		} else {
			if intervals[i][0] < intervals[j][0] {
				intervals[j][0] = intervals[i][0]
			}
			if intervals[i][1] > intervals[j][1] {
				intervals[j][1] = intervals[i][1]
			}
		}
	}
	return intervals[:j+1]
}

func sort1(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	left, right := 0, len(intervals)-1
	r := rand.Int() % len(intervals)

	intervals[r], intervals[right] = intervals[right], intervals[r]

	for i := 0; i < right; i++ {
		if intervals[i][0] < intervals[right][0] {
			intervals[i], intervals[left] = intervals[left], intervals[i]
			left++
		}
	}
	intervals[left], intervals[right] = intervals[right], intervals[left]
	sort1(intervals[:left])
	sort1(intervals[left+1:])
	return intervals
}

func rotate(matrix [][]int) {
	len := len(matrix)
	//对角反转
	for row := 0; row < len; row++ {
		for col := 0; col <= row; col++ {
			matrix[row][col] = matrix[col][row]
		}
	}
	//垂直翻转
	for i := 0; i < len; i++ {
		for j := 0; j < (len+1)/2; j++ {
			matrix[i][j] = matrix[i][len-j-1]
		}
	}
}
