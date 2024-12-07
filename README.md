# go-iforest

## Installation

```bash
go get github.com/narumiruna/go-iforest
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/narumiruna/go-iforest/pkg/iforest"
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
```
