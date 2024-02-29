package leetcodelearn

import (
	"fmt"
	"testing"
)

func TestDinnerPlateStacks(t *testing.T) {

	dp := Constructor(1)

	dp.Push(1)
	dp.Push(2)

	fmt.Println(dp.PopAtStack(1))
	fmt.Println(dp.Pop())
	dp.Push(1)
	dp.Push(2)
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())

}
