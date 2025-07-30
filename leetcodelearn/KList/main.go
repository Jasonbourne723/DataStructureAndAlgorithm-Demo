package main

import (
	"fmt"
)

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	n := reverseKGroup(node1, 2)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	dummyHead := &ListNode{0, head}
	lastNode := dummyHead
	for isMatch(lastNode, k) {
		var h *ListNode
		node := lastNode.Next
		for i := 0; i < k; i++ {
			h, node, node.Next = node, node.Next, h
		}
		lastNode.Next = h
		lastNode = node
	}
	return dummyHead.Next
}

func isMatch(last *ListNode, k int) bool {
	node := last
	for i := 0; i < k; i++ {
		if node == nil {
			return false
		}
		node = node.Next
		if node == nil {
			return false
		}
	}
	return true
}
