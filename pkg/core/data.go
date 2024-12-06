package core

import (
	"math"
	"math/rand/v2"
)

func sample(data [][]float64, sampleSize int) [][]float64 {
	if len(data) <= sampleSize {
		return data
	}

	perm := rand.Perm(len(data))
	sampledData := make([][]float64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sampledData[i] = data[perm[i]]
	}
	return sampledData
}

func maxValue(data []float64) float64 {
	maxValue := math.Inf(-1)
	for _, v := range data {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func minValue(data []float64) float64 {
	minValue := math.Inf(1)
	for _, v := range data {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

func slice(data [][]float64, attributeIndex int) []float64 {
	slicedData := make([]float64, len(data))
	for i, row := range data {
		slicedData[i] = row[attributeIndex]
	}
	return slicedData
}
