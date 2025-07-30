package sort

// 最大数是 4
func BaseSort(array []int) []int {

	temp := [5]int{}

	for _, item := range array {
		temp[item]++
	}

	i := 0
	for j := 0; j < 5; j++ {
		for temp[j] > 0 {
			array[i] = j
			i++
			temp[j]--
		}
	}
	return array
}
