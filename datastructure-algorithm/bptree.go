package DataStructureAndAlgorithm

import (
	"fmt"
	"sync"
)

type BPItem struct {
	Key int64
	Val interface{}
}

type BPNode struct {
	MaxKey int64
	Nodes  []*BPNode
	Items  []BPItem
	Next   *BPNode
}

type BPTree struct {
	mutex sync.RWMutex
	root  *BPNode
	width int
	halfw int
}

func NewBPTree(width int) *BPTree {
	if width < 3 {
		width = 3
	}
	var bt = &BPTree{}
	bt.root = NewLeafNode(width)
	bt.width = width
	bt.halfw = (bt.width + 1) / 2
	return bt
}

func NewLeafNode(width int) *BPNode {
	var node = &BPNode{}
	node.Items = make([]BPItem, width+1)
	node.Items = node.Items[0:0]
	return node
}

func NewIndexNode(width int) *BPNode {
	var node = &BPNode{}
	node.Nodes = make([]*BPNode, width+1)
	node.Nodes = node.Nodes[0:0]
	return node
}

func (t *BPTree) Set(key int64, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.setValue(nil, t.root, key, value)
}

func (t *BPTree) setValue(parent *BPNode, node *BPNode, key int64, value interface{}) {
	for i := 0; i < len(node.Nodes); i++ {
		if key <= node.Nodes[i].MaxKey || i == len(node.Nodes)-1 {
			t.setValue(node, node.Nodes[i], key, value)
			break
		}
	}

	//叶子结点，添加数据
	if len(node.Nodes) < 1 {
		node.setValue(key, value)
	}

	//结点分裂
	node_new := t.splitNode(node)
	if node_new != nil {
		//若父结点不存在，则创建一个父节点
		if parent == nil {
			parent = NewIndexNode(t.width)
			parent.addChild(nil, node)
			t.root = parent
		}
		//添加结点到父亲结点
		parent.addChild(node, node_new)
	}
}

func (n *BPNode) addChild(node *BPNode, nodeNew *BPNode) {

	if node == nil {
		n.Nodes = append(n.Nodes, nodeNew)
	} else {

		for i := 0; i < len(n.Nodes); i++ {
			if n.Nodes[i] == node {
				temp := append(n.Nodes[:i+1], nodeNew)
				if i+1 < len(n.Nodes) {
					n.Nodes = append(temp, n.Nodes[i+1:]...)
				} else {
					n.Nodes = temp
				}
				break
			}
		}
	}
}

func (t *BPTree) splitNode(node *BPNode) *BPNode {
	if len(node.Nodes) > t.width {
		//创建新结点
		halfw := t.width/2 + 1
		node2 := NewIndexNode(t.width)
		node2.Nodes = append(node2.Nodes, node.Nodes[halfw:len(node.Nodes)]...)
		node2.MaxKey = node2.Nodes[len(node2.Nodes)-1].MaxKey

		//修改原结点数据
		node.Nodes = node.Nodes[0:halfw]
		node.MaxKey = node.Nodes[len(node.Nodes)-1].MaxKey

		return node2
	} else if len(node.Items) > t.width {
		//创建新结点
		halfw := t.width/2 + 1
		node2 := NewLeafNode(t.width)
		node2.Items = append(node2.Items, node.Items[halfw:len(node.Items)]...)
		node2.MaxKey = node2.Items[len(node2.Items)-1].Key
		node2.Next = node.Next

		//修改原结点数据
		node.Next = node2
		node.Items = node.Items[0:halfw]
		node.MaxKey = node.Items[len(node.Items)-1].Key

		return node2
	}

	return nil
}

func (node *BPNode) setValue(key int64, value interface{}) {
	item := BPItem{key, value}
	num := len(node.Items)
	if num < 1 {
		node.Items = append(node.Items, item)
		node.MaxKey = item.Key
		return
	} else if key < node.Items[0].Key {
		node.Items = append([]BPItem{item}, node.Items...)
		return
	} else if key > node.Items[num-1].Key {
		node.Items = append(node.Items, item)
		node.MaxKey = item.Key
		return
	}

	for i := 0; i < num; i++ {
		if node.Items[i].Key > key {
			node.Items = append(node.Items, BPItem{})
			copy(node.Items[i+1:], node.Items[i:])
			node.Items[i] = item
			return
		} else if node.Items[i].Key == key {
			node.Items[i] = item
			return
		}
	}
}

func (t *BPTree) Get(key int64) interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	node := t.root
	for i := 0; i < len(node.Nodes); i++ {
		if key <= node.Nodes[i].MaxKey {
			node = node.Nodes[i]
			i = 0
		}
	}

	//没有到达叶子结点
	if len(node.Nodes) > 0 {
		return nil
	}

	for i := 0; i < len(node.Items); i++ {
		if node.Items[i].Key == key {
			return node.Items[i].Val
		}
	}
	return nil
}

func (t *BPTree) List() {
	t.mutex.RLocker().Lock()
	defer t.mutex.RLocker().Unlock()

	node := t.root

	for {
		if len(node.Nodes) > 0 {
			node = node.Nodes[0]
		} else {
			break
		}
	}

	for {
		for i := 0; i < len(node.Items); i++ {
			fmt.Printf("%d\t", node.Items[i].Key)
		}
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}

}
