package DataStructureAndAlgorithm

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	stack := NewStack[int]()
	stack.Push(104)
	stack.Push(103)
	stack.Push(102)
	stack.Push(101)

	for {
		if data, ok := stack.Pop(); ok {
			fmt.Printf("data: %v\n", data)
		} else {
			break
		}
	}
}
