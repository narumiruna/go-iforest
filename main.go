package main

import (
	"fmt"

	"github.com/narumiruna/go-isolation-forest/pkg/core"
	"github.com/narumiruna/go-isolation-forest/pkg/types"
)

func main() {

	// x := types.RandomMatrix(1000, 2)
	// y := types.RandomMatrix(10, 2).MulScalar(1000).AddScalar(10)
	x := types.Matrix{
		{1.1, 2.2},
		{1.1, 2.2},
		{1.1, 2.2},
		{1.1, 2.2},
		{1.1, 2.2},
		{10, 20},
	}
	fmt.Println(x)
	fmt.Println(x.Shape())

	model := core.NewIsolationForest()
	model.Fit(x)
	for _, tree := range model.Trees {
		fmt.Println(tree)
	}

	scores := model.Score(x)
	fmt.Println(scores)

	predicts := model.Predict(x)
	fmt.Println(predicts)
}
