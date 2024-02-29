package leetcodelearn

type DinnerPlates struct {
	Capacity   int
	Stacks     []Stack
	FreeStacks []int
}

func (d *DinnerPlates) AddFreeStacks(index int) {

	if len(d.FreeStacks) == 0 {
		d.FreeStacks = append(d.FreeStacks, index)
		return
	}

	if index > d.FreeStacks[len(d.FreeStacks)-1] {
		d.FreeStacks = append(d.FreeStacks, index)
		return
	}

	for i, v := range d.FreeStacks {
		if index < v {
			temp := append(d.FreeStacks[0:i], index)
			d.FreeStacks = append(temp, d.FreeStacks[i:]...)

		}
	}

}

func Constructor(capacity int) DinnerPlates {
	plates := DinnerPlates{Capacity: capacity}

	plates.Stacks = make([]Stack, 0, 5)
	plates.FreeStacks = make([]int, 0)
	return plates
}

func (d *DinnerPlates) Push(val int) {

	l := len(d.Stacks)

	if len(d.FreeStacks) > 0 {
		for i := range d.FreeStacks {

			if l == 0 || i >= l {
				break
			}

			if ok := d.Stacks[i].Push(val); ok {
				return
			}
		}
	}

	if l == 0 || len(d.Stacks[l-1].data) == d.Capacity {
		stack := *NewStack(d.Capacity)
		stack.Push(val)
		d.Stacks = append(d.Stacks, stack)
	} else {
		d.Stacks[l-1].data = append(d.Stacks[l-1].data, val)
	}

}

func (d *DinnerPlates) Pop() int {
	for i := len(d.Stacks) - 1; i >= 0; i-- {
		if ok, val := d.Stacks[i].Pop(); ok {
			d.resize()
			return val
		}
	}
	return -1
}

func (d *DinnerPlates) PopAtStack(index int) int {

	if index >= len(d.Stacks) {
		return -1
	}

	if ok, val := d.Stacks[index].Pop(); ok {
		if index == len(d.Stacks)-1 {
			d.resize()
		}
		d.AddFreeStacks(index)
		return val
	} else {
		return -1
	}
}

func (d *DinnerPlates) resize() {
	l := len(d.Stacks)

	if len(d.Stacks[l-1].data) == 0 {
		d.Stacks = d.Stacks[:l-1]
	}

}

type Stack struct {
	capacity int
	data     []int
}

func NewStack(capacity int) *Stack {
	stack := &Stack{capacity: capacity}
	stack.data = make([]int, 0, capacity)
	return stack
}

func (s *Stack) Push(val int) bool {
	if len(s.data) == s.capacity {
		return false
	} else {
		s.data = append(s.data, val)
		return true
	}
}

func (s *Stack) Pop() (bool, int) {
	if len(s.data) > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return true, val
	} else {
		return false, 0
	}
}
