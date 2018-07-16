package order

import (
	"go/build"

	"github.com/cohadar/gopkgorder/graph"
	"golang.org/x/tools/go/buildutil"
)

// creates a new graph from build context.
// example: GetGraph(&build.Default)
func GetGraph(context *build.Context) (g graph.Graph, err error) {
	g = graph.NewGraph()
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
				g.AddEdge(path, absPath(imp))
			}
			// range bp.TestImports not used to avoid cycles
		}
	})
	return
}
