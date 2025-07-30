package tree

import (
	"fmt"
)

func NewBinaryTree() *BinaryTree {
	tree := new(BinaryTree)
	return tree
}

type BinaryTree struct {
	root *Node
}

type Node struct {
	item   int
	count  int
	left   *Node
	right  *Node
	height int8
}

func GetHeight(node *Node) int8 {
	if node == nil {
		return -1
	}
	return node.height
}

func (t *BinaryTree) Insert1(item int) {
	if t.root == nil {
		t.root = &Node{
			item:   item,
			count:  1,
			height: 0,
		}
		return
	}
	insert1(t.root, item)
}

func (b *BinaryTree) Insert(item int) {
	b.root = insert(b.root, item)
}

func insert(node *Node, item int) *Node {
	if node == nil {
		return &Node{
			item:   item,
			count:  1,
			height: 0,
		}
	}
	if item > node.item {
		node.right = insert(node.right, item)
		if GetHeight(node.right)-GetHeight(node.left) > 1 {
			if item > node.right.item {
				node = singleRotateRight(node)
			} else {
				node = doubleRotateRight(node)
			}
		}
	} else if item < node.item {
		node.left = insert(node.left, item)
		if GetHeight(node.left)-GetHeight(node.right) > 1 {
			if item > node.left.item {
				node = doubleRotateLeft(node)
			} else {
				node = singleRotateLeft(node)
			}
		}
	} else {
		node.count++
	}
	node.height = maxInt8(GetHeight(node.left), GetHeight(node.right)) + 1
	return node
}

func insert1(node *Node, item int) {

	if item > node.item {
		if node.right == nil {
			node.right = &Node{
				item:   item,
				count:  1,
				height: 0,
			}
		} else {
			insert1(node.right, item)
		}
	} else if item < node.item {
		if node.left == nil {
			node.left = &Node{
				item:   item,
				count:  1,
				height: 0,
			}
		} else {
			insert1(node.left, item)
		}
	} else {
		node.count++
	}
	node.height = maxInt8(GetHeight(node.left), GetHeight(node.right)) + 1
}

func maxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func singleRotateLeft(n *Node) *Node {
	n1 := n.left
	n.left = n1.right
	n1.right = n
	n1.height = maxInt8(GetHeight(n1.left), GetHeight(n1.right)) + 1
	n.height = maxInt8(GetHeight(n.left), GetHeight(n.right)) + 1
	return n1
}

func doubleRotateLeft(n *Node) *Node {

	n.left = singleRotateRight(n.left)
	return singleRotateLeft(n)
}

func doubleRotateRight(n *Node) *Node {
	n.right = singleRotateLeft(n.right)
	return singleRotateRight(n)
}

func singleRotateRight(n *Node) *Node {
	n1 := n.right
	n.right = n1.left
	n1.left = n
	n1.height = maxInt8(GetHeight(n1.left), GetHeight(n1.right)) + 1
	n.height = maxInt8(GetHeight(n.left), GetHeight(n.right)) + 1
	return n1
}

func (t *BinaryTree) Print() {
	print(t.root)
}

func print(n *Node) {
	if n == nil {
		return
	}
	for i := 0; i < n.count; i++ {
		fmt.Printf("n.item: %v,height:%v\n", n.item, n.height)
	}
	print(n.left)
	print(n.right)
}

func (t *BinaryTree) Get(item int) bool {
	return get(t.root, item)
}

func get(n *Node, item int) bool {
	if n == nil {
		return false
	}
	if n.item == item {
		if n.count > 0 {
			return true
		} else {
			return false
		}
	} else if item > n.item {
		return get(n.right, item)
	} else {
		return get(n.left, item)
	}
}
