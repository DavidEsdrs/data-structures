package graphs

import "github.com/DavidEsdrs/ads/sets"

func (g *Graph[T]) LargestComponent() int {
	visited := sets.NewSet[T, bool]()
	largest := 0
	for vertex := range g.GetMap() {
		if !visited.Has(vertex) {
			currentSize := 0
			g.DepthFirst(vertex, func(current T) {
				visited.Add(current, true)
				currentSize++
			})
			if currentSize > largest {
				largest = currentSize
			}
		}
	}
	return largest
}
