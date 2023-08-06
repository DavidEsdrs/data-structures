package queue

type Queue[T any] struct {
	_nodes []T
	Size   uint
}

func (q *Queue[T]) Push(n T) {
	q._nodes = append(q._nodes, n)
	q.Size++
}

func (q *Queue[T]) Shift() T {
	var zeroValue T
	if q.Size == 0 {
		return zeroValue
	}
	first := q._nodes[0]
	q._nodes = q._nodes[1:]
	q.Size--
	return first
}
