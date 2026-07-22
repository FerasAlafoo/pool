package pool

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	lastItem := len(s.items) - 1
	removedValue := s.items[lastItem]
	s.items = s.items[:lastItem]
	return removedValue
}
