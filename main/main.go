package main

import (
	"fmt"
	"sort"

	"github.com/cohadar/gopkgorder"
)

func links(row []string) string {
	ret := ""
	for _, pkg := range row {
		// TODO: add dependencies in hower text and maybe dependency count as superscript
		ret += fmt.Sprintf("[%s](https://golang.org/pkg/%s \"%s\") ", pkg, pkg, pkg)
	}
	return ret
}

func main() {
	rows, err := gopkgorder.GetResult()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		sort.Strings(row)
		fmt.Println(links(row))
	}
}
