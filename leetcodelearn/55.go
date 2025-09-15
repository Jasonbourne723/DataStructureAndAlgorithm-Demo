package leetcodelearn

func CanJump(nums []int) bool {
	var f int
	for i, v := range nums {
		if i < f {
			return false
		}
		t := i + v
		if t > f {
			f = t
		}
	}
	return true
}
