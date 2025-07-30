package DataStructureAndAlgorithm

import "testing"

func TestHeapSort(t *testing.T) {
	array := []int{3, 2, 1, 4, 5}
	result := HeapSort(array)
	if result[0] != 1 {
		t.Errorf("HeapSort failed, expected %v, got %v", 1, result[0])
	}
	if result[1] != 2 {
		t.Errorf("HeapSort failed, expected %v, got %v", 2, result[1])
	}
	if result[2] != 3 {
		t.Errorf("HeapSort failed, expected %v, got %v", 3, result[2])
	}
	if result[3] != 4 {
		t.Errorf("HeapSort failed, expected %v, got %v", 4, result[3])
	}
	if result[4] != 5 {
		t.Errorf("HeapSort failed, expected %v, got %v", 5, result[4])
	}
}
