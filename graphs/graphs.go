package graphs

import (
	"github.com/DavidEsdrs/ads/queue"
	"github.com/DavidEsdrs/ads/sets"
	"github.com/DavidEsdrs/ads/stacks"
)

type Graph[T comparable] struct {
	Map map[T][]T
}

func NewGraph[T comparable]() Graph[T] {
	return Graph[T]{
		Map: make(map[T][]T),
	}
}

func (g *Graph[T]) Insert(key T) {
	if _, exists := g.Map[key]; !exists {
		g.Map[key] = []T{}
	}
}

func (g *Graph[T]) CreateEdge(key T, to ...T) {
	if _, exists := g.Map[key]; !exists {
		g.Map[key] = []T{}
	}
	for _, str := range to {
		_, exists := g.Map[str]
		if !exists {
			g.Insert(str)
		}
		g.Map[key] = append(g.Map[key], str)
	}
}

func (g *Graph[T]) CreateEdgeBidirectional(key T, to ...T) {
	if _, exists := g.Map[key]; !exists {
		g.Map[key] = []T{}
	}
	for _, str := range to {
		_, exists := g.Map[str]
		if !exists {
			g.Insert(str)
		}
		g.Map[key] = append(g.Map[key], str)
		g.Map[str] = append(g.Map[str], key)
	}
}

func (g *Graph[T]) GetNeighbours(src T) (neighbours []T, exists bool) {
	neighbours, exists = g.Map[src]
	return neighbours, exists
}

type Callback[T comparable] func(current T)

func (g *Graph[T]) DepthFirst(start T, cb Callback[T]) {
	stack := stacks.Stack[T]{}
	_, exists := g.Map[start]
	if !exists {
		return
	}
	set := sets.NewSet[T, bool]()
	stack.Push(start)
	for stack.Size > 0 {
		current := stack.Pop()
		if !set.Has(current) {
			cb(current)
			set.Add(current, true)
			edges := g.Map[current]
			for _, n := range edges {
				stack.Push(n)
			}
		}
	}
}

func (g *Graph[T]) DepthFirstRecursive(start T, cb Callback[T]) {
	cb(start)
	node := g.Map[start]
	for _, neighbor := range node {
		g.DepthFirstRecursive(neighbor, cb)
	}
}

func (g *Graph[T]) BreadthFirst(start T, cb Callback[T]) {
	queue := queue.Queue[T]{}
	_, exists := g.Map[start]
	if !exists {
		panic("Root not found")
	}
	set := sets.NewSet[T, bool]()
	queue.Push(start)
	for queue.Size > 0 {
		current := queue.Shift()
		if !set.Has(current) {
			cb(current)
			edges := g.Map[current]
			for _, neighbor := range edges {
				queue.Push(neighbor)
			}
		}
	}
}

func DepthFirst[T comparable](g *Graph[T], start T) {
	stack := stacks.Stack[T]{}
	stack.Push(start)
	visited := sets.NewSet[T, bool]()
	for stack.Size > 0 {
		current := stack.Pop()
		if !visited.Has(current) {
			visited.Add(current, true)
			println(current)
			neighbours := g.Map[current]
			for _, n := range neighbours {
				stack.Push(n)
			}
		}
	}
}
