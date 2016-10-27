package main

import (
	"fmt"
	"go/build"
	"sort"
	"strings"

	"github.com/cohadar/gopkgorder"
)

func dependencyText(deps []string) (ret string) {
	sort.Strings(deps)
	ret = strings.Join(deps, ", ")
	if ret == "" {
		ret = "none"
	}
	return
}

func links(graph gopkgorder.Graph, row []string) string {
	ret := ""
	for _, pkg := range row {
		deps := graph.GetDependencies(pkg)
		ret += fmt.Sprintf("[%s{%d}](https://golang.org/pkg/%s \"%s\")\n", pkg, len(deps), pkg, dependencyText(deps))
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
	for _, row := range rows {
		sort.Strings(row)
		fmt.Println(links(graph, row))
	}
}
