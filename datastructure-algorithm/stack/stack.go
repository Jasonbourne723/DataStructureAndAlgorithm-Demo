package stack

type Stack[T any] struct {
	array     []T
	lastIndex int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		array:     make([]T, 0),
		lastIndex: -1,
	}
}

func (s *Stack[T]) Pop() (data T, isExist bool) {

	l := len(s.array)
	if l == 0 {
		return data, false
	}
	data = s.array[l-1]
	s.array = s.array[:l-1]
	return data, true
}

func (s *Stack[T]) Push(data T) {
	s.array = append(s.array, data)
}
