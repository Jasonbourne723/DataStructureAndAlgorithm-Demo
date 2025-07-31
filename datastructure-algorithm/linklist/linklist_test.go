package linklist

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {

	linkList := &LinkList[int]{}
	linkList.Append(10)
	linkList.Append(101)
	linkList.Append(102)
	linkList.Append(130)
	linkList.Append(150)
	linkList.Append(11)
	linkList.Append(9)
	linkList.Append(29)
	linkList.Append(10)

	l := linkList.Find(func(i int) bool {
		return i == 10
	})
	for i, item := range l {
		fmt.Printf("%d,%d\n", i, item)
	}

	val, exist := linkList.First(func(i int) bool { return i == 151 })
	fmt.Printf("x: %d, exist: %v\n", val, exist)

	fmt.Println("----------------")

	linkList.Reversal()
	fmt.Println("----------------")
	//linkList.print()

	linkList.Append(45)
	linkList.print()
}
