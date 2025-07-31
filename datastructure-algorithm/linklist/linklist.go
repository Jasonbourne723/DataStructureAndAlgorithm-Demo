package linklist

import "fmt"

type LinkList[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
}

type ListNode[T any] struct {
	data T
	next *ListNode[T]
}

func (l *LinkList[T]) Append(data T) {
	if l.head == nil {
		l.head = &ListNode[T]{
			data: data,
		}
		l.tail = l.head
	} else {
		node := ListNode[T]{
			data: data,
		}
		l.tail.next = &node
		l.tail = l.tail.next
	}
}

func (l *LinkList[T]) Find(filterFunc func(T) bool) []T {
	var list []T
	t := l.head
	for t != nil {
		if filterFunc(t.data) {
			list = append(list, t.data)
		}
		t = t.next
	}
	return list
}

func (l *LinkList[T]) First(filterFunc func(T) bool) (T, bool) {

	t := l.head
	for t != nil {
		if filterFunc(t.data) {
			return t.data, true
		}
		t = t.next
	}
	var data T
	return data, false
}

// Reversal 头插法反转
func (l *LinkList[T]) Reversal() {

	if l.head == nil {
		return
	}
	var newHead *ListNode[T]
	var tail *ListNode[T]
	t := l.head
	for t != nil {
		node := ListNode[T]{
			data: t.data,
			next: newHead,
		}
		if tail == nil {
			tail = &node
		}
		newHead = &node
		t = t.next
	}
	l.head = newHead
	l.tail = tail
}

func (l *LinkList[T]) print() {
	t := l.head

	for t != nil {

		fmt.Println(t.data)

		t = t.next
	}
}
