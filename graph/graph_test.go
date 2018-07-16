package graph_test

import (
	"fmt"

	"github.com/cohadar/gopkgorder/graph"
)

func ExampleGraph_AddNode() {
	g := graph.NewGraph()
	g.AddNode("Alfa")
	g.AddNode("Bravo")
	g.AddNode("Charlie")
	fmt.Println(g.GetNodes())
	// Output: [Alfa Bravo Charlie]
}

func ExampleGraph_Size() {
	g := graph.NewGraph()
	g.AddNode("Alfa")
	g.AddNode("Bravo")
	g.AddNode("Charlie")
	fmt.Println(g.Size())
	// Output: 3
}

func ExampleGraph_HasNode() {
	g := graph.NewGraph()
	g.AddNode("Alfa")
	g.AddNode("Bravo")
	g.AddNode("Charlie")
	fmt.Println(g.HasNode("Bravo"), g.HasNode("Delta"))
	// Output: true false
}

func ExampleGraph_AddEdge() {
	g := graph.NewGraph()
	g.AddEdge("Alfa", "Bravo")
	fmt.Println(g.HasEdge("Alfa", "Bravo"), g.HasEdge("Bravo", "Alfa"))
	// Output: true false
}

func ExampleGraph_GetRoots() {
	g := graph.NewGraph()
	g.AddEdge("Charlie", "Alfa")
	g.AddEdge("Delta", "Alfa")
	g.AddEdge("Echo", "Charlie")
	g.AddEdge("Echo", "Bravo")
	fmt.Println(g.GetRoots())
	// Output: [Alfa Bravo]
}

func ExampleGraph_RemoveNode() {
	g := graph.NewGraph()
	g.AddEdge("Bravo", "Alfa")
	g.AddEdge("Charlie", "Bravo")
	g.RemoveNode("Bravo")
	fmt.Println(g.GetLeafs())
	fmt.Println(g.HasEdge("Charlie", "Bravo"))
	// Output:
	// [Alfa Charlie]
	// false
}

func ExampleGraph_GetTargets() {
	g := graph.NewGraph()
	g.AddEdge("Alfa", "Bravo")
	g.AddEdge("Alfa", "Charlie")
	g.AddEdge("Alfa", "Delta")
	g.AddEdge("Echo", "Alfa")
	fmt.Println(g.GetTargets("Alfa"))
	// Output: [Bravo Charlie Delta]
}
