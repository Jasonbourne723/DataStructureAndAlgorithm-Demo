package leetcodelearn

func threeSum(nums []int) [][]int {

	m := make(map[int]int, len(nums))
	nums2 := make([]int, 0, 10)
	r := [][]int{}

	j := 0
	for _, v := range nums {
		if _, exist := m[v]; !exist {
			m[v] = j
			nums2 = append(nums2, v)
			j++
		}
	}

	for i := 0; i < len(nums2)-2; i++ {
		for j := i + 1; j < len(nums2)-1; j++ {
			sum := (0 - (nums2[i] + nums2[j]))
			if value, exist := m[sum]; exist {
				if value > j {
					r = append(r, []int{nums2[i], nums2[j], nums2[value]})
				}
			}
		}
	}

	return r

}
