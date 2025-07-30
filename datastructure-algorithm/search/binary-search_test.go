package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	dataStore := NewDataStore([]int{1, 3, 4, 6, 7, 9, 10, 12, 35, 47, 56})

	if index, ok := dataStore.BinarySearch(6); ok {
		fmt.Printf("index1: %v\n", index)
	}

	if index, ok := dataStore.BinarySearch(77); ok {
		fmt.Printf("index2: %v\n", index)
	}

}
