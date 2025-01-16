package leetcodelearn

func merge1(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}

	p := m + n - 1

	for n > 0 || m > 0 {
		if n == 0 || (m > 0 && nums1[m-1] > nums2[n-1]) {
			nums1[p] = nums1[m-1]
			p--
			m--
		} else {
			nums1[p] = nums2[n-1]
			p--
			n--
		}
	}
	return

}
