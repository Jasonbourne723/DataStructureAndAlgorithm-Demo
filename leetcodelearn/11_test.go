package leetcodelearn

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {

	height := [...]int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxArea(height[:])

	fmt.Printf("area: %v\n", area)
}
