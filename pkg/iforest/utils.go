package iforest

import (
	"math"
	"math/rand"
)

func Clone(m [][]float64) [][]float64 {
	clone := make([][]float64, len(m))
	for i, row := range m {
		clone[i] = make([]float64, len(row))
		copy(clone[i], row)
	}
	return clone
}

func Sample(m [][]float64, sampleSize int) [][]float64 {
	if sampleSize <= 0 {
		panic("sampleSize must be greater than 0")
	}

	if len(m) <= sampleSize {
		return m
	}

	perm := rand.Perm(len(m))
	sampled := make([][]float64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sampled[i] = m[perm[i]]
	}
	return sampled
}

func Column(m [][]float64, j int) []float64 {
	column := make([]float64, len(m))
	for i, row := range m {
		column[i] = row[j]
	}
	return column
}

func MinMax(slice []float64) (float64, float64) {
	min, max := math.Inf(1), math.Inf(-1)
	for _, v := range slice {
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
	m := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			m[i][j] = rand.Float64()
		}
	}
	return m
}

func AddScalar(m [][]float64, scalar float64) [][]float64 {
	clone := Clone(m)
	for i, row := range clone {
		for j := range row {
			clone[i][j] += scalar
		}
	}
	return clone
}
