package sort

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2
	return merge(MergeSort(array[:mid]), MergeSort(array[mid:]))
}

func merge(left, right []int) []int {

	n := len(left)
	m := len(right)
	i, j := 0, 0
	result := make([]int, 0, n+m)

	for i < n || j < m {
		if i == n || (j < m && left[i] > right[j]) {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	return result
}
