package graphs

import (
	"github.com/DavidEsdrs/ads/queue"
	"github.com/DavidEsdrs/ads/sets"
	"github.com/DavidEsdrs/ads/stacks"
)

func (g *Graph[T]) HasPath(source T, dest T) bool {
	stack := stacks.Stack[T]{}
	_, exists := g.Map[source]
	if !exists {
		return false
	}
	visited := sets.NewSet[T, bool]()
	stack.Push(source)
	for stack.Size > 0 {
		current := stack.Pop()
		if !visited.Has(current) {
			if current == dest {
				return true
			}
			visited.Add(current, true)
			neighbours := g.Map[current]
			for _, n := range neighbours {
				stack.Push(n)
			}
		}
	}
	return false
}

type CompareFunction[T comparable] func(a T, b T) bool

func (g *Graph[T]) HasPathWithCompare(source T, dest T, compareFunction CompareFunction[T]) bool {
	stack := stacks.Stack[T]{}
	_, exists := g.Map[source]
	if !exists {
		return false
	}
	visited := sets.NewSet[T, bool]()
	stack.Push(source)
	for stack.Size > 0 {
		current := stack.Pop()
		if !visited.Has(current) {
			if compareFunction(current, dest) {
				return true
			}
			visited.Add(current, true)
			neighbours := g.Map[current]
			for _, n := range neighbours {
				stack.Push(n)
			}
		}
	}
	return false
}

func (g *Graph[T]) HasPathRecursive(source T, dest T) bool {
	if source == dest {
		return true
	}
	neighbours := g.Map[source]
	for _, n := range neighbours {
		if g.HasPathRecursive(n, dest) {
			return true
		}
	}
	return false
}

func (g *Graph[T]) HasPathBreadthFirst(source T, dest T) bool {
	queue := queue.Queue[T]{}
	_, exists := g.Map[source]
	if !exists {
		return false
	}
	queue.Push(source)
	for queue.Size > 0 {
		current := queue.Shift()
		if current == dest {
			return true
		}
		neighbours := g.Map[current]
		for _, n := range neighbours {
			queue.Push(n)
		}
	}
	return false
}
