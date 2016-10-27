package gopkgorder

import (
	"go/build"

	"golang.org/x/tools/go/buildutil"
)

// graph of package import dependencies
type Graph map[string]map[string]bool

// creates a new graph from build context.
// example: GetGraph(&build.Default)
func GetGraph(context *build.Context) (graph Graph, err error) {
	graph = make(Graph)
	defer func() {
		if r := recover(); r != nil {
			err = recover().(error)
		}
		return
	}()
	buildutil.ForEachPackage(context, func(path string, err error) {
		if err != nil {
			panic(err)
		}
		bp, err := context.Import(path, "", 0)
		if err != nil {
			if _, ok := err.(*build.NoGoError); ok {
				// empty directory is valid package
			} else {
				panic(err)
			}
		}
		absPath := func(path string) string {
			if bp2, _ := context.Import(path, bp.Dir, build.FindOnly); bp2 != nil {
				return bp2.ImportPath
			}
			return path
		}
		if bp != nil {
			for _, imp := range bp.Imports {
				graph.addEdge(path, absPath(imp))
			}
			// range bp.TestImports not used to avoid cycles
		}
	})
	return
}

func (g Graph) addEdge(from, to string) {
	if to == "C" {
		return // "C" is fake package
	}
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
