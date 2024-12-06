package main

import (
	"fmt"

	"github.com/narumiruna/go-isolation-forest/pkg/core"
	"github.com/narumiruna/go-isolation-forest/pkg/types"
)

func main() {
	dim := 2
	x := types.RandomMatrix(1000, dim)
	y := types.RandomMatrix(10, dim).AddScalar(0.5)

	model := core.NewIsolationForest()
	model.Fit(x)

	scores := model.Score(y)
	fmt.Println(scores)

	predicts := model.Predict(y)
	fmt.Println(predicts)
}
