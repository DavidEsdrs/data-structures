package graphs

import (
	"github.com/DavidEsdrs/ads/sets"
)

func (g *Graph[T]) GetMap() map[T][]T {
	return g.Map
}

func (g *Graph[T]) CountComponents() int {
	visited := sets.NewSet[T, bool]()
	components := 0
	for vertex := range g.GetMap() {
		if !visited.Has(vertex) {
			g.DepthFirst(vertex, func(current T) {
				visited.Add(current, true)
			})
			components++
		}
	}
	return components
}

func NumberOfComponents[T comparable](graph *Graph[T]) int {
	components := 0
	visited := sets.NewSet[T, bool]()
	for vertex := range graph.GetMap() {
		if depthFirst(graph, vertex, &visited) {
			components++
		}
	}
	return components
}

func depthFirst[T comparable](graph *Graph[T], current T, visited *sets.Set[T, bool]) bool {
	if visited.Has(current) {
		return false
	}
	visited.Add(current, true)
	neighbours := graph.Map[current]
	for _, n := range neighbours {
		depthFirst(graph, n, visited)
	}
	return true
}
