package graph

import (
	"sort"
)

// Graph represent directed acyclic graph of strings
type Graph map[string]map[string]bool

// NewGraph makes a new Graph
func NewGraph() Graph {
	return make(Graph)
}

// Size returns number of nodes
func (g Graph) Size() int {
	return len(g)
}

// AddNode adds a new unconnected node into graph if it does not exist
func (g Graph) AddNode(node string) {
	if edges, ok := g[node]; !ok {
		edges = make(map[string]bool)
		g[node] = edges
	}
}

// GetNodes returns a sorted lice with node names
func (g Graph) GetNodes() []string {
	var nodes []string
	for node := range g {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)
	return nodes
}

// HasNode return true if node exists in graph
func (g Graph) HasNode(node string) bool {
	_, ok := g[node]
	return ok
}

// AddEdge adds a from --> to, if the nodes do not exist, it creates them
func (g Graph) AddEdge(from, to string) {
	g.AddNode(from)
	g.AddNode(to)
	g[to][from] = true
}

// HasEdge return true if edge from --> to exists
func (g Graph) HasEdge(from, to string) bool {
	if set, ok := g[to]; ok {
		return set[from]
	}
	return false
}

// GetLeafs returns nodes with no dependencies (sorted)
func (g Graph) GetRoots() []string {
	var leafs []string
	for name := range g {
		if len(g.GetTargets(name)) == 0 {
			leafs = append(leafs, name)
		}
	}
	sort.Strings(leafs)
	return leafs
}

// RemoveNode removes node from graph
func (g Graph) RemoveNode(node string) {
	delete(g, node)
	for _, set := range g {
		delete(set, node)
	}
}

// GetTargets returns all nodes pointed by node from
func (g Graph) GetTargets(from string) []string {
	var ret []string
	for to, set := range g {
		if set[from] {
			ret = append(ret, to)
		}
	}
	sort.Strings(ret)
	return ret
}
