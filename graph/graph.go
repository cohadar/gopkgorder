package graph

import "strings"

// Graph represent package dependencies
type Graph map[string]map[string]bool

// AddEdge from dependant to used package
func (g Graph) AddEdge(from, to string) {
	// if to == "C" {
	// 	return // "C" is fake package
	// }
	g.addNode(from)
	g.addNode(to)
	g[from][to] = true
}

func (g Graph) addNode(node string) {
	edges := g[node]
	if edges == nil {
		edges = make(map[string]bool)
		g[node] = edges
	}
}

func NewGraph() Graph {
	return make(Graph)
}

func (g Graph) GetDependencies(root string) (ret []string) {
	m := make(map[string]bool)
	for from, subMap := range g {
		if strings.HasPrefix(from, root) {
			for name, _ := range subMap {
				if !strings.HasPrefix(name, root) {
					m[name] = true
				}
			}
		}
	}
	for name, _ := range m {
		ret = append(ret, name)
	}
	return
}
