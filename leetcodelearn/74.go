package leetcodelearn

func SearchMatrix(matrix [][]int, target int) bool {

	l_x := len(matrix)
	l_y := len(matrix[0])

	return binarySearch1(0, l_x*l_y-1, matrix, target)
}

func binarySearch1(left int, right int, matrix [][]int, target int) bool {

	if left > right {
		return false
	}
	if getValue(left, matrix) > target || getValue(right, matrix) < target {
		return false
	}
	if getValue(left, matrix) == target || getValue(right, matrix) == target {
		return true
	}
	middle := (right-left)/2 + left
	if getValue(middle, matrix) == target {
		return true
	}
	if getValue(middle, matrix) < target {
		return binarySearch1(middle+1, right-1, matrix, target)
	} else {
		return binarySearch1(left+1, middle-1, matrix, target)
	}

}

func getValue(index int, matrix [][]int) int {

	i := index / len(matrix[0])

	j := index % len(matrix[0])

	return matrix[i][j]

}
