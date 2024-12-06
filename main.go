package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/narumiruna/go-isolation-forest/pkg/iforest"
	"gonum.org/v1/gonum/mat"
)

func RandomMatrix(n, dim int, bias float64) *mat.Dense {
	data := make([]float64, n*dim)
	for i := range data {
		data[i] = rand.Float64() + bias
	}
	return mat.NewDense(n, dim, data)
}

func main() {
	dim := 2
	x := RandomMatrix(1000, dim, 0)
	y := RandomMatrix(10, dim, 0.5)

	model := iforest.New()
	model.Fit(x)

	scores := model.Score(y)
	fmt.Println(scores)

	predicts := model.Predict(y)
	fmt.Println(predicts)
}
