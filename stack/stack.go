package stack

type Stack[T any] struct {
	_nodes []*T
	Size   uint
}

func (s *Stack[T]) Push(n *T) {
	s._nodes = append(s._nodes, n)
	s.Size++
}

func (s *Stack[T]) Pop() *T {
	last := s._nodes[len(s._nodes)-1]
	s._nodes = s._nodes[:len(s._nodes)-1]
	s.Size--
	return last
}

func (s *Stack[T]) Peek() *T {
	return s._nodes[len(s._nodes)-1]
}
