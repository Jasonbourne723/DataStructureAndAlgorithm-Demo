package leetcodelearn

import (
	"fmt"
	"testing"
)

func Test6(t *testing.T) {
	s := "AB"
	s1 := convert(s, 1)
	fmt.Printf("s1: %v\n", s1)
}

func TestFindErrorNums(t *testing.T) {

	ans := findErrorNums([]int{3, 2, 3, 4, 6, 5})
	fmt.Printf("ans: %v\n", ans)
}

func TestRemoveDuplicates(t *testing.T) {
	a := removeDuplicates([]int{1, 1, 1, 2, 2, 3})
	fmt.Printf("a: %v\n", a)
}
