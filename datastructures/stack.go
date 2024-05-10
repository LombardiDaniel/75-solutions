package datastructures

type Stack struct {
	arr []int
}

// LIFO
func NewStack() Stack {
	return Stack{
		arr: []int{},
	}
}

func (s *Stack) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() int {
	ret := s.arr[len(s.arr) - 1]
	s.arr = s.arr[:len(s.arr) - 1]

	return ret
}
