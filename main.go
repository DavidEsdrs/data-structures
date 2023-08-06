package main

import "github.com/DavidEsdrs/ads/graphs"

type User struct {
	Name string
}

func main() {
	graph := graphs.NewGraph[string]()
	graph.Insert("a")
	graph.Insert("b")
	graph.Insert("c")
	graph.Insert("d")
	graph.Insert("e")
	graph.Insert("f")

	graph.CreateEdge("a", "b", "c")
	graph.CreateEdge("b", "d")
	graph.CreateEdge("c", "e")
	graph.CreateEdge("d", "f")

	list := [][2]User{
		{{Name: "Ana"}, {Name: "Bruna"}},
		{{Name: "Ana"}, {Name: "Carla"}},
		{{Name: "Bruna"}, {Name: "Carla"}},
		{{Name: "Bruna"}, {Name: "David"}},
		{{Name: "Carla"}, {Name: "David"}},
		{{Name: "David"}, {Name: "Ernando"}},
		{{Name: "Ernando"}, {Name: "Fernando"}},
		{{Name: "Fernando"}, {Name: "Gabriel"}},
		{{Name: "Gabriel"}, {Name: "Heitor"}},
		{{Name: "Heitor"}, {Name: "Ernando"}},
	}

	g := graphs.ToGraph(list)

	// println(g.HasPath(User{Name: "David"}, User{Name: "Ernando"}))

	g.DepthFirst(User{Name: "Carla"}, func(current User) {
		if current.Name != "Carla" {
			println("Carla is connected to", current.Name)
		}
	})
}
