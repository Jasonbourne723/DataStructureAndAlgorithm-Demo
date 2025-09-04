package leetcodelearn

import "reflect"

func Subsets(nums []int) [][]int {
	result := [][]int{}
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {

		if !isExist(result, path) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)

		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}

func isExist(result [][]int, path []int) bool {

	for _, item := range result {
		if reflect.DeepEqual(item, path) {
			return true
		}
	}
	return false
}
