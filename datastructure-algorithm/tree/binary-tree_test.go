package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	tree := NewBinaryTree()

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)

	tree.Print()

	fmt.Printf("tree.Get(50): %v\n", tree.Get(50))
	fmt.Printf("tree.Get(90): %v\n", tree.Get(90))
}
