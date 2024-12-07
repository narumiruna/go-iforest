package iforest

import (
	"math"
	"math/rand"
)

type Vector []float64

func ZeroVector(size int) Vector {
	return make(Vector, size)
}

func RandomVector(size int) Vector {
	vector := make(Vector, size)
	for i := range vector {
		vector[i] = rand.Float64()
	}
	return vector
}

func (v Vector) Clone() Vector {
	clone := make(Vector, len(v))
	copy(clone, v)
	return clone
}

func (v Vector) Size() int {
	return len(v)
}

func (v Vector) ZerosLike() Vector {
	return ZeroVector(len(v))
}

func (v Vector) AddScalarInPlace(scalar float64) Vector {
	for i := range v {
		v[i] += scalar
	}
	return v
}

func (v Vector) AddScalar(scalar float64) Vector {
	return v.Clone().AddScalarInPlace(scalar)
}

func (v Vector) MulScalarInPlace(scalar float64) Vector {
	for i := range v {
		v[i] *= scalar
	}
	return v
}

func (v Vector) MulScalar(scalar float64) Vector {
	return v.Clone().MulScalarInPlace(scalar)
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

func (v Vector) MinMax() (float64, float64) {
	minValue, maxValue := math.Inf(1), math.Inf(-1)
	for _, value := range v {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}
	return minValue, maxValue
}
