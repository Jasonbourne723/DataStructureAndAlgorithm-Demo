package DataStructureAndAlgorithm

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {

	a := []int{1, 2, 3, 4}
	a = reverse(a)
	fmt.Printf("a: %v\n", a)
}
