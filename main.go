package main

import (
	"fmt"

	"github.com/narumiruna/go-isolation-forest/pkg/core"
)

func main() {
	data := [][]float64{
		{1.1, 2.2},
		{3.3, 4.4},
		{5.5, 6.6},
		{7.7, 8.8},
		{9.9, 10.10},
	}

	model := core.NewIsolationForest()
	model.Fit(data)

	for _, tree := range model.Trees {
		fmt.Println(tree)
	}

}
