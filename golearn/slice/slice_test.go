package slice_test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	s := []int{1, 2, 3, 4}

	Add(s)

	fmt.Printf("s: %v\n", s)

}

func Add(s []int) {

	s[1] = 3
}
