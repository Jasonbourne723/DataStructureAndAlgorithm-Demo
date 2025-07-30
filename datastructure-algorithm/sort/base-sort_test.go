package sort

import (
	"fmt"
	"testing"
)

func TestBaseSort(t *testing.T) {
	arrary := [10]int{4, 3, 3, 2, 0, 0, 1, 2, 3, 4}
	arrary1 := BaseSort(arrary[:])
	fmt.Printf("arrary1: %v\n", arrary1)
}
