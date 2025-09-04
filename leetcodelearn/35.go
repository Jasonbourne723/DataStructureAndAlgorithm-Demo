package leetcodelearn

func SearchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left > right {
		return left
	}
	if right < 0 {
		return 0
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}

	if left == right {
		if target > nums[left] {
			return left + 1
		} else {
			return left
		}
	}

	middle := (right-left)/2 + left
	if nums[middle] == target {
		return middle
	}

	if target > middle {
		return binarySearch(nums, middle+1, right, target)
	} else {
		return binarySearch(nums, left, middle-1, target)
	}

}
