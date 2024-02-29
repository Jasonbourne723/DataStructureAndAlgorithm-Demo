package leetcodelearn

type DinnerPlates struct {
	Capacity   int
	Stacks     []Stack
	FreeStacks []int
}

func (this *DinnerPlates) AddFreeStacks(index int) {

	if len(this.FreeStacks) == 0 {
		this.FreeStacks = append(this.FreeStacks, index)
		return
	}

	if index > this.FreeStacks[len(this.FreeStacks)-1] {
		this.FreeStacks = append(this.FreeStacks, index)
		return
	}

	for i, v := range this.FreeStacks {
		if index < v {
			temp := append(this.FreeStacks[0:i], index)
			this.FreeStacks = append(temp, this.FreeStacks[i:]...)

		}
	}

}

func Constructor(capacity int) DinnerPlates {
	plates := DinnerPlates{Capacity: capacity}

	plates.Stacks = make([]Stack, 0, 5)
	plates.FreeStacks = make([]int, 0)
	return plates
}

func (this *DinnerPlates) Push(val int) {

	l := len(this.Stacks)

	if len(this.FreeStacks) > 0 {
		for i, _ := range this.FreeStacks {

			if l == 0 || i >= l {
				break
			}

			if ok := this.Stacks[i].Push(val); ok {
				return
			}
		}
	}

	if l == 0 || len(this.Stacks[l-1].data) == this.Capacity {
		stack := *NewStack(this.Capacity)
		stack.Push(val)
		this.Stacks = append(this.Stacks, stack)
	} else {
		this.Stacks[l-1].data = append(this.Stacks[l-1].data, val)
	}

}

func (this *DinnerPlates) Pop() int {
	for i := len(this.Stacks) - 1; i >= 0; i-- {
		if ok, val := this.Stacks[i].Pop(); ok {
			this.resize()
			return val
		}
	}
	return -1
}

func (this *DinnerPlates) PopAtStack(index int) int {

	if index >= len(this.Stacks) {
		return -1
	}

	if ok, val := this.Stacks[index].Pop(); ok {
		if index == len(this.Stacks)-1 {
			this.resize()
		}
		this.AddFreeStacks(index)
		return val
	} else {
		return -1
	}
}

func (this *DinnerPlates) resize() {
	l := len(this.Stacks)

	if len(this.Stacks[l-1].data) == 0 {
		this.Stacks = this.Stacks[:l-1]
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
