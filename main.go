package main

import (
	"fmt"

	"github.com/narumiruna/go-isolation-forest/pkg/iforest"
)

func main() {
	dim := 2
	x := iforest.RandomMatrix(1000, dim)
	y := iforest.RandomMatrix(10, dim).AddScalar(0.5)

	model := iforest.NewWithOptions(
		iforest.Options{
			NumTrees:      100,
			SampleSize:    256,
			DetectionType: "threshold",
			Proportion:    0.6,
		},
	)

	model.Fit(x)

	scores := model.Score(y)
	fmt.Println(scores)

	predictions := model.Predict(y)
	fmt.Println(predictions)

	importances := model.FeatureImportance(y[0])
	fmt.Println(importances)
}
