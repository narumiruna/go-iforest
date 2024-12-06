package main

import (
	"fmt"

	"github.com/narumiruna/go-isolation-forest/pkg/core"
)

func main() {
	data := [][]float64{
		{1.1, 2.2},
		{1.1, 2.2},
		{1.1, 2.2},
		{1.1, 2.2},
		{1.1, 2.2},
		{10, 20},
	}

	model := core.NewIsolationForest()
	model.Fit(data)

	for _, tree := range model.Trees {
		fmt.Println(tree)
	}

	scores := model.Score(data)
	fmt.Println(scores)

	predicts := model.Predict(data)
	fmt.Println(predicts)
}
