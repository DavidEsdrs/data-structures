package queue

type Queue[T any] struct {
	_nodes []*T
	Size   uint
}

func (q *Queue[T]) Unshift(n *T) {
	q._nodes = append(q._nodes, n)
	q.Size++
}

func (q *Queue[T]) Shift() *T {
	if q.Size == 0 {
		return nil
	}
	first := q._nodes[0]
	q._nodes = q._nodes[1:]
	q.Size--
	return first
}
