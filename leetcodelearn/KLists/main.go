/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

func main() {

	// Example usage
	lists := []*ListNode{} // Initialize your linked lists here

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3

	lists = append(lists, node1)

	node11 := &ListNode{Val: 1}
	node21 := &ListNode{Val: 3}
	node31 := &ListNode{Val: 4}
	node11.Next = node21
	node21.Next = node31

	lists = append(lists, node11)

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 6}
	node4.Next = node5
	lists = append(lists, node4)

	mergedList := mergeKLists(lists)
	for mergedList != nil {
		fmt.Printf("%d -> ", mergedList.Val)
		mergedList = mergedList.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type heap struct {
	data []*ListNode
}

func NewHeap(cap int) *heap {
	return &heap{
		data: make([]*ListNode, 0, cap),
	}
}

func (h *heap) add(node *ListNode) {
	if len(h.data) == 0 {
		h.data = append(h.data, node)
	} else {
		h.data = append(h.data, node)
		t := len(h.data) - 1
		for t > 0 {
			if h.data[t].Val < h.data[(t-1)/2].Val {
				h.data[t], h.data[(t-1)/2] = h.data[(t-1)/2], h.data[t]
				t = (t - 1) / 2
			} else {
				break
			}
		}
	}
}

func (h *heap) top() *int {
	if len(h.data) == 0 {
		return nil
	}
	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	t := 0
	for {
		t1 := t
		left := 2*t + 1
		right := 2*t + 2
		if left < len(h.data) && h.data[left].Val < h.data[t1].Val {
			t1 = left
		}
		if right < len(h.data) && h.data[right].Val < h.data[t1].Val {
			t1 = right
		}
		if t1 == t {
			break
		}
		h.data[t], h.data[t1] = h.data[t1], h.data[t]
	}
	if top.Next != nil {
		h.add(top.Next)
	}
	return &top.Val
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var node *ListNode
	h := NewHeap(len(lists))
	f := 0

	for i, _ := range lists {
		if lists[i] != nil {
			h.add(lists[i])
		}
	}
	for {
		v := h.top()
		if v == nil {
			break
		}
		if f == 0 {
			node = &ListNode{
				Val: *v,
			}
			head = node
			f += 1
		} else {
			n := &ListNode{
				Val: *v,
			}
			node.Next = n
			node = node.Next
		}
	}

	return head
}
