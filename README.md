# go-iforest

Go-iforest is a Go implementation of the Isolation Forest algorithm for anomaly detection. It identifies anomalies by isolating observations in random decision trees.

[![GoDoc](https://pkg.go.dev/badge/github.com/narumiruna/go-iforest.svg)](https://pkg.go.dev/github.com/narumiruna/go-iforest)
[![Go Report Card](https://goreportcard.com/badge/github.com/narumiruna/go-iforest)](https://goreportcard.com/report/github.com/narumiruna/go-iforest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
