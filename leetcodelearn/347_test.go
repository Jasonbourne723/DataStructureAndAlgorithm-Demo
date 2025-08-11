package leetcodelearn

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	ar := []int{1, 1, 1, 2, 2, 3}
	result := TopKFrequent(ar, 2)
	fmt.Print(result)
}
