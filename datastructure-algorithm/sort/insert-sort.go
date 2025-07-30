package sort

func InsertSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		j := i
		for j > 0 && array[j] < array[j-1] {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}
	}
	return array
}
