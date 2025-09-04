package leetcodelearn

import (
	"fmt"
	"testing"
)

func TestMinWindow(test *testing.T) {
	s := "ADOBECODEBANC"
	t := "ABC"
	r := MinWindow(s, t)
	fmt.Println(r)
}
