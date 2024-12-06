package types

import (
	"math"
	"math/rand/v2"
)

type Vector []float64

func ZeroVector(n int) Vector {
	return make(Vector, n)
}

func RandomVector(n int) Vector {
	v := make(Vector, n)
	for i := range v {
		v[i] = rand.Float64()
	}
	return v
}

func (v Vector) Shape() int {
	return len(v)
}

func (v Vector) AddScalar(scalar float64) Vector {
	o := make(Vector, len(v))
	for i := range v {
		o[i] = v[i] + scalar
	}
	return o
}

func (v Vector) MulScalar(scalar float64) Vector {
	o := make(Vector, len(v))
	for i := range v {
		o[i] = v[i] * scalar
	}
	return o
}

func (v Vector) Max() float64 {
	maxValue := math.Inf(-1)
	for _, value := range v {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func (v Vector) Min() float64 {
	minValue := math.Inf(1)
	for _, value := range v {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}
