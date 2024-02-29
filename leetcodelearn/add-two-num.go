package leetcodelearn

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
*/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {

	var tail *ListNode
	var carry = 0

	for l1 != nil || l2 != nil {
		t1, t2 := 0, 0

		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		}

		sum := t1 + t2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
