package leetcodelearn

import (
	"fmt"
	"testing"
)

func TestAtoi(t *testing.T) {

	a := myAtoi("-5-")

	fmt.Printf("a: %v\n", a)
}
func TestThreeSumClosest(t *testing.T) {
	r := threeSumClosest([]int{2, 5, 6, 7}, 1)
	fmt.Printf("r: %v\n", r)
}
