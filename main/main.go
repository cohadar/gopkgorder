package main

import (
	"fmt"

	"github.com/cohadar/gopkgorder"
)

func main() {
	rows, err := gopkgorder.GetResult()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		fmt.Println(row)	
	}
}
