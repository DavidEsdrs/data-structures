package tree

import (
	"github.com/DavidEsdrs/ads/queue"
	"github.com/DavidEsdrs/ads/stack"
)

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type Tree[T any] struct {
	Root *Node[T]
}

func (t *Tree[Y]) InDepthTraversal() []*Node[Y] {
	stack := stack.Stack[Node[Y]]{}
	res := []*Node[Y]{}
	stack.Push(t.Root)
	for stack.Size > 0 {
		current := stack.Pop()
		res = append(res, current)
		println(current.Value)
		if current.Right != nil {
			stack.Push(current.Right)
		}
		if current.Left != nil {
			stack.Push(current.Left)
		}
	}
	return res
}

func (t *Tree[Y]) InDepthTraversalRecursive() []*Node[Y] {
	res := recurse(t.Root)
	return res
}

func recurse[T any](n *Node[T]) []*Node[T] {
	if n == nil {
		return []*Node[T]{}
	}
	left := recurse(n.Left)
	right := recurse(n.Right)
	res := []*Node[T]{}
	res = append(res, left...)
	res = append(res, right...)
	return res
}

func (t *Tree[Y]) BreadthFirstTraversal() []*Node[Y] {
	queue := queue.Queue[Node[Y]]{}
	res := []*Node[Y]{}
	queue.Unshift(t.Root)
	for queue.Size > 0 {
		current := queue.Shift()
		res = append(res, current)
		println(current.Value)
		if current.Left != nil {
			queue.Unshift(current.Left)
		}
		if current.Right != nil {
			queue.Unshift(current.Right)
		}
	}
	return res
}

func Execute() {
	a := Node[int]{Value: 3}
	b := Node[int]{Value: 11}
	c := Node[int]{Value: 4}
	d := Node[int]{Value: 4}
	e := Node[int]{Value: 2}
	f := Node[int]{Value: 1}

	a.Left = &b
	a.Right = &c

	b.Left = &d
	b.Right = &e

	c.Right = &f

	// ========================================
	g := Node[int]{Value: 5}

	h := Node[int]{Value: 11}
	i := Node[int]{Value: 3}

	j := Node[int]{Value: 4}
	k := Node[int]{Value: 15}

	l := Node[int]{Value: 12}

	m := Node[int]{Value: 1}

	g.Left = &h
	g.Right = &i

	h.Left = &j
	h.Right = &k

	i.Right = &l

	j.Right = &m

	t := Tree[int]{Root: &g}

	println(Height(t.Root))
}
