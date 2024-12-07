package iforest

import (
	"math"
	"math/rand"
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

func (v Vector) Clone() Vector {
	o := make(Vector, len(v))
	copy(o, v)
	return o
}

func (v Vector) Size() int {
	return len(v)
}

func (v Vector) ZerosLike() Vector {
	return ZeroVector(len(v))
}

func (v Vector) AddScalar_(scalar float64) Vector {
	for i := range v {
		v[i] += scalar
	}
	return v
}

func (v Vector) AddScalar(scalar float64) Vector {
	return v.Clone().AddScalar_(scalar)
}

func (v Vector) MulScalar_(scalar float64) Vector {
	for i := range v {
		v[i] *= scalar
	}
	return v
}

func (v Vector) MulScalar(scalar float64) Vector {
	return v.Clone().MulScalar_(scalar)
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
