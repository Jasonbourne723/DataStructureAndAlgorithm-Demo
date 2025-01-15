package DataStructureAndAlgorithm

import "fmt"

type LinkList struct {
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	data int
	next *LinkNode
	priv *LinkNode
}

func (l *LinkList) Insert(data int) {

	if l.head == nil {
		l.head = &LinkNode{
			data: data,
			next: nil,
			priv: nil,
		}
		l.tail = l.head
		return
	}
	t := l.head
	i := 0
	for t != nil {
		if i == 0 {
			if data < t.data {
				node := &LinkNode{
					data: data,
					next: t,
					priv: nil,
				}
				l.head.priv = node
				l.head = node
				return
			}
		}
		if data < t.data {
			node := &LinkNode{
				data: data,
				next: t,
				priv: t.priv,
			}
			t.priv.next = node
			t.priv = node
			return
		}
		t = t.next
		i++
	}
	node := &LinkNode{
		data: data,
		next: nil,
		priv: l.tail,
	}
	l.tail.next = node
	l.tail = node
}

func (l *LinkList) Print() {
	n := l.head
	for n != nil {
		fmt.Printf("n.data: %v\n", n.data)
		n = n.next
	}
}

func (l *LinkList) ReversePrint() {
	n := l.tail
	for n != nil {
		fmt.Printf("n.data: %v\n", n.data)
		n = n.priv
	}
}
