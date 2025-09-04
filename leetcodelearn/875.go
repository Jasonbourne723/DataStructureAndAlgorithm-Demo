package leetcodelearn

func MinEatingSpeed(piles []int, h int) int {

	var max int
	for _, item := range piles {
		if item > max {
			max = item
		}
	}
	curr := BinarySearch(piles, 1, max, h)

	return curr

}

func BinarySearch(piles []int, min int, max int, h int) int {
	// 每h吃最少个时，需要的h最大
	t_min := eat(piles, min)
	if h >= t_min {
		return min
	}
	//// 每h吃最多时，需要的h最小
	t_max := eat(piles, max)
	if h < t_max {
		// 吃不完
		return -1
	}

	if min >= max {
		return max
	}

	m := (max-min)/2 + min
	count_m := eat(piles, m)

	if count_m <= h {
		return BinarySearch(piles, min+1, m, h)
	} else {
		return BinarySearch(piles, m+1, max, h)
	}

}

func eat(piles []int, k int) int {

	var total int
	for _, item := range piles {
		total += gettime(item, k)
	}
	return total

}

func gettime(count int, k int) int {
	t := count / k
	if count%k > 0 {
		t += 1
	}
	return t
}
