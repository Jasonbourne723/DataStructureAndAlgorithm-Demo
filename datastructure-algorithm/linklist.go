package DataStructureAndAlgorithm

import "fmt"

type LinkList struct {
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	data int
	next *LinkNode
}

func (l *LinkList) Insert(data int) {
	if l.head == nil {
		l.head = &LinkNode{
			data: data,
		}
		l.tail = l.head
	} else {

		if data > l.head.data {
			node := &LinkNode{
				data: data,
				next: l.head,
			}
			l.head = node
		} else {
			var node = l.head
			for {
				last := node
				node = node.next
				if node == nil {
					last.next = &LinkNode{
						data: data,
					}
					break
				} else {
					if data > node.data {
						tempNode := &LinkNode{
							data: data,
							next: node,
						}
						last.next = tempNode
						break
					}
				}
			}
		}
	}
}

func (l *LinkList) PrintAll() {
	var node = l.head

	for {
		if node != nil {
			fmt.Printf("node.data: %v\n", node.data)
			node = node.next
		} else {
			break
		}
	}

}
