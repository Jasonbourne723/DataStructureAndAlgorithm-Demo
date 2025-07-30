package linklist

import "testing"

func TestLinkList(t *testing.T) {

	linkList := &LinkList{}

	linkList.Insert(10)
	linkList.Insert(101)
	linkList.Insert(102)
	linkList.Insert(130)
	linkList.Insert(150)
	linkList.Insert(11)
	linkList.Insert(9)
	linkList.Insert(29)

	linkList.Print()
	linkList.ReversePrint()
}
