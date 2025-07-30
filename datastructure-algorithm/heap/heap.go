package heap

import (
	"errors"
)

// 大顶堆
type Heap struct {
	cap       int32
	lastIndex int32
	array     []int32
}

func NewHeap(cap int32) *Heap {
	return &Heap{
		cap:       cap,
		lastIndex: 0,
		array:     make([]int32, cap+1),
	}
}

// 插入
func (h *Heap) Insert(data int32) error {
	if h.lastIndex > h.cap {
		return errors.New("heap is full")
	}
	h.lastIndex++
	h.array[h.lastIndex] = data
	currIndex := h.lastIndex
	for {
		if currIndex == 1 {
			break
		}

		i := currIndex / 2
		if h.array[currIndex] > h.array[i] {
			h.swap(currIndex, i)
			currIndex = i
		} else {
			break
		}
	}
	return nil
}

// 删除堆顶
func (h *Heap) DeleteTop() int32 {
	if h.lastIndex == 1 {
		return -1
	}
	top := h.array[1]
	h.array[1] = h.array[h.lastIndex]
	h.lastIndex--
	i := 1
	for {
		maxPos := i
		if i*2 < int(h.lastIndex) && h.array[i] < h.array[i*2] {
			maxPos = i * 2
		}
		if i*2+1 < int(h.lastIndex) && h.array[maxPos] < h.array[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		h.swap(int32(i), int32(maxPos))
		i = maxPos
	}
	return top
}

// 交换数组元素
func (h *Heap) swap(aIndex int32, bIndex int32) {
	h.array[aIndex] ^= h.array[bIndex]
	h.array[bIndex] ^= h.array[aIndex]
	h.array[aIndex] ^= h.array[bIndex]
}
