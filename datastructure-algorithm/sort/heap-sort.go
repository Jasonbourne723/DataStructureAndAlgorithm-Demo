package DataStructureAndAlgorithm

import "fmt"

// HeapSort sorts an array using the heap sort algorithm
func HeapSort(array []int) []int {
	n := len(array)

	for i := 1; i < n; i++ {
		k := i
		j := (i - 1) / 2
		for j >= 0 && array[j] < array[k] {
			array[k], array[j] = array[j], array[k]
			if j == 0 {
				break
			}
			k = j
			j = (j - 1) / 2
		}
	}

	fmt.Println(array)

	for i := 0; i < n; i++ {
		array[0], array[n-1-i] = array[n-1-i], array[0]
		j := 0
		for {
			k := 2*j + 1
			if k >= n-1-i {
				break
			}
			if k+1 < n-1-i && array[k] < array[k+1] {
				k++
			}
			if array[j] < array[k] {
				array[j], array[k] = array[k], array[j]
				j = k
			} else {
				break
			}
		}
	}

	fmt.Println(array)

	return array
}
