package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap(20)

	heap.Insert(21)
	heap.Insert(32)
	heap.Insert(42)
	heap.Insert(25)
	heap.Insert(27)
	heap.Insert(82)
	heap.Insert(92)
	heap.Insert(22)
	heap.Insert(234)
	heap.Insert(24)

	fmt.Printf(" %v\n", heap.DeleteTop())
	fmt.Printf(" %v\n", heap.DeleteTop())
	fmt.Printf(" %v\n", heap.DeleteTop())
	fmt.Printf(" %v\n", heap.DeleteTop())
	fmt.Printf(" %v\n", heap.DeleteTop())
}
