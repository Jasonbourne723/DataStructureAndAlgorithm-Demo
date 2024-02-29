package leetcodelearn

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {

	var m = make(map[int]int, 10)

	for i, v := range nums {
		var other = target - v

		if otherVal, ok := m[other]; ok {
			return []int{otherVal, i}
		} else {
			m[v] = i
		}
	}
	return nil
}
