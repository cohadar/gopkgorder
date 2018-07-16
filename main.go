package main

import (
	"fmt"
	"go/build"
	"log"
	"runtime"
	"sort"
	"strings"

	"github.com/cohadar/gopkgorder/graph"
	"github.com/cohadar/gopkgorder/order"
)

const (
	HEADER = `## Go standard library package order, from least to most dependent.
Hover to see dependency list. Number in curly braces is just a number of direct imports. Each row is one level in topsorted graph.
`
	LINK = "[%s{%d}](https://golang.org/pkg/%s \"%s\")\n"
)

func dependencyText(deps []string) (ret string) {
	sort.Strings(deps)
	ret = strings.Join(deps, ", ")
	if ret == "" {
		ret = "nil"
	}
	return
}

func links(g graph.Graph, row []string) string {
	ret := ""
	for _, pkg := range row {
		deps := g.GetTargets(pkg)
		ret += fmt.Sprintf(LINK, pkg, len(deps), pkg, dependencyText(deps))
	}
	return ret
}

func prune(g graph.Graph) {
	nodes := g.GetNodes()
	for _, node := range nodes {
		root := strings.Split(node, "/")[0]
		// remove github.com and golang.org base packages
		if strings.ContainsRune(root, '.') {
			g.RemoveNode(node)
		}
		// tools are not library package
		if root == "cmd" {
			g.RemoveNode(node)
		}
	}
}

func main() {
	g, err := order.GetGraph(&build.Default)
	if err != nil {
		panic(err)
	}
	prune(g)
	var rows [][]string
	// extract roots until nothing in graph
	for g.Size() > 0 {
		roots := g.GetRoots()
		for _, root := range roots {
			g.RemoveNode(root)
		}
		if len(roots) == 0 && g.Size() != 0 {
			log.Fatal("cycle in graph", g)
		}
		rows = append(rows, roots)
	}
	// build graph again because ExtractOrderedPackages is destructive
	g, err = order.GetGraph(&build.Default)
	if err != nil {
		panic(err)
	}
	fmt.Print(HEADER)
	fmt.Printf("Go Version: %s\n\n", runtime.Version())
	for _, row := range rows {
		fmt.Println(links(g, row))
	}
}
