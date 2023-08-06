package graphs

import "github.com/DavidEsdrs/ads/sets"

func ToDigraph[T comparable](edges [][2]T) *Graph[T] {
	g := NewGraph[T]()
	for _, connections := range edges {
		a, b := connections[0], connections[1]
		g.Insert(a)
		g.Insert(b)

		g.CreateEdge(a, b)
		g.CreateEdge(b, a)
	}
	return &g
}

func ToGraph[T comparable](edges [][2]T) *Graph[T] {
	g := NewGraph[T]()
	for _, connections := range edges {
		a, b := connections[0], connections[1]
		g.Insert(a)
		g.Insert(b)

		g.CreateEdge(a, b)
	}
	return &g
}

func UndirectedPath(edges [][2]string, src string, dst string) bool {
	set := sets.NewSet[string, bool]()
	graph := ToDigraph(edges)
	return hasPath(graph, src, dst, &set)
}

func hasPath[T comparable](graph *Graph[T], src T, dst T, visited *sets.Set[T, bool]) bool {
	if visited.Has(src) {
		return false
	}
	if src == dst {
		return true
	}
	visited.Add(src, true)
	neighbours, _ := graph.GetNeighbours(src)
	for _, n := range neighbours {
		if hasPath(graph, n, dst, visited) {
			return true
		}
	}
	return false
}
