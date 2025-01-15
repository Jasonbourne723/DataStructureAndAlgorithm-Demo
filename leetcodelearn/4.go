package leetcodelearn

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)

	m1 := (l1 - 1) / 2
	m2 := (l2 - 1) / 2

	if l1 == 0 {
		if l2 == 0 {
			return 0
		}
		return middle(nums2)
	}

	if l2 == 0 {
		return middle(nums1)
	}

	if l1 == l2 && l1 == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}

	for {
		if m1+1 < l1 && nums1[m1+1] < nums2[m2] {
			m1++
			m2--
		} else if m2+1 < l2 && nums2[m2+1] < nums1[m1] {
			m2++
			m1--
		} else {
			if (l1+l2)%2 == 0 {
				var left int
				var right int

				if m1 == -1 {
					left = nums2[m2]
				} else if m2 == -1 {
					left = nums1[m1]
				} else {
					left = max(nums1[m1], nums2[m2])
				}

				if m1 >= l1-1 {
					right = nums2[m2+1]
				} else if m2 >= l2-1 {
					right = nums1[m1+1]
				} else {
					right = min(nums1[m1+1], nums2[m2+1])
				}

				return (float64(left) + float64(right)) / 2
			} else {
				if m1 == -1 {
					return float64(nums2[m2])
				} else if m2 == -1 {
					return float64(nums1[m1])
				} else {
					return float64(max(nums1[m1], nums2[m2]))
				}
			}
		}
	}

}

func middle(nums []int) float64 {
	l := len(nums)

	if l == 1 {
		return float64(nums[0])
	}

	if len(nums)%2 == 0 {
		return float64(nums[(l-1)/2]+nums[(l-1)/2+1]) / 2
	} else {
		return float64(nums[(l-1)/2])
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
