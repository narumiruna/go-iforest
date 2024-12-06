package main

import (
	"fmt"

	"github.com/narumiruna/go-isolation-forest/pkg/core"
	"github.com/narumiruna/go-isolation-forest/pkg/types"
)

func main() {

	x := types.RandomMatrix(1000, 2)
	y := types.RandomMatrix(10, 2).MulScalar(1).AddScalar(1)
	fmt.Println("y[0] min:", y[0].Min())

	// x := types.Matrix{
	// 	{1.1, 2.2},
	// 	{1.1, 2.2},
	// 	{1.1, 2.2},
	// 	{1.1, 2.2},
	// 	{1.1, 2.2},
	// 	{10, 20},
	// }
	// fmt.Println(x)
	// fmt.Println(x.Shape())

	model := core.NewIsolationForest()
	model.Fit(x)
	for _, tree := range model.Trees {
		fmt.Println(tree)
	}

	scores := model.Score(y)
	fmt.Println(scores)

	predicts := model.Predict(y)
	fmt.Println(predicts)
}
