package main

import (
	"fmt"

	"github.com/narumiruna/go-iforest/pkg/iforest"
)

// main function demonstrates the usage of the iForest package for anomaly detection.
// It generates random data, fits the model, scores the data, makes predictions, and calculates feature importances.
func main() {
	dim := 2
	x := iforest.RandomMatrix(1000, dim)
	y := iforest.RandomMatrix(10, dim)
	y = iforest.AddScalar(y, 0.5)

	model := iforest.NewWithOptions(
		iforest.Options{
			NumTrees:      100,
			SampleSize:    256,
			DetectionType: "threshold",
			Threshold:     0.6,
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
