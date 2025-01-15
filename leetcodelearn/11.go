package leetcodelearn

func maxArea(height []int) int {

	if len(height) < 2 {
		return 0
	}

	l := 0
	r := len(height) - 1

	area := min(height[l], height[r]) * (r - l)

	for l != r {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		temp := min(height[l], height[r]) * (r - l)
		area = max(temp, area)
	}

	return area

}
