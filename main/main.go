package main

import (
	"fmt"
	"go/build"
	"sort"
	"strings"

	"github.com/cohadar/gopkgorder"
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

func links(graph gopkgorder.Graph, row []string) string {
	ret := ""
	for _, pkg := range row {
		deps := graph.GetDependencies(pkg)
		ret += fmt.Sprintf(LINK, pkg, len(deps), pkg, dependencyText(deps))
	}
	return ret
}

// TODO: generate README.md with template
func main() {
	graph, err := gopkgorder.GetGraph(&build.Default)
	if err != nil {
		panic(err)
	}
	rows, err := graph.GetOrderedPackages()
	if err != nil {
		panic(err)
	}
	// build graph again because GetOrderedPackages is destructive
	graph, err = gopkgorder.GetGraph(&build.Default)
	if err != nil {
		panic(err)
	}
	fmt.Println(HEADER)
	for _, row := range rows {
		sort.Strings(row)
		fmt.Println(links(graph, row))
	}
}
