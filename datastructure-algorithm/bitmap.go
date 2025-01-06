package DataStructureAndAlgorithm

func NewBitMap(size int32) *BitMap {

	b := new(BitMap)
	b.array = make([]int, (size+31)/32)
	return b
}

type BitMap struct {
	array []int
}

func (b *BitMap) Set(num int32) {

	b.array[num/32] |= 1 << (num % 32)
}

func (b *BitMap) Remove(num int32) {
	b.array[num/32] &= ^(1 << (num % 32))
}

func (b *BitMap) Test(num int32) bool {
	return (b.array[num/32] & (1 << (num % 32))) != 0
}
