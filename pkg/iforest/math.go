package iforest

import (
	"math"
	"math/rand"
)

func Sample(matrix [][]float64, size int) [][]float64 {
	if size <= 0 {
		panic("size must be greater than 0")
	}

	if len(matrix) <= size {
		return matrix
	}

	perm := rand.Perm(len(matrix))
	sampled := make([][]float64, size)
	for i := 0; i < size; i++ {
		sampled[i] = matrix[perm[i]]
	}
	return sampled
}

func Column(matrix [][]float64, columnIndex int) []float64 {
	column := make([]float64, len(matrix))
	for i, row := range matrix {
		column[i] = row[columnIndex]
	}
	return column
}

func MinMax(floats []float64) (float64, float64) {
	min, max := math.Inf(1), math.Inf(-1)
	for _, v := range floats {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func RandomMatrix(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Float64()
		}
	}
	return matrix
}

func AddScalar(matrix [][]float64, scalar float64) [][]float64 {
	outputMatrix := make([][]float64, len(matrix))
	for i, row := range matrix {
		outputMatrix[i] = make([]float64, len(row))
		for j := range row {
			outputMatrix[i][j] = matrix[i][j] + scalar
		}
	}
	return outputMatrix
}
