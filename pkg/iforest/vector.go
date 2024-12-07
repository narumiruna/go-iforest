// Package iforest provides implementations of isolation forest algorithms for anomaly detection.
package iforest

import (
	"math"
	"math/rand"
)

// Vector represents a slice of float64 values.
type Vector []float64

// ZeroVector returns a zero-initialized Vector of the given size.
func ZeroVector(size int) Vector {
	return make(Vector, size)
}

// RandomVector returns a Vector of the given size with random float64 values.
func RandomVector(size int) Vector {
	vector := make(Vector, size)
	for i := range vector {
		vector[i] = rand.Float64()
	}
	return vector
}

// Clone creates a copy of the Vector.
func (v Vector) Clone() Vector {
	clone := make(Vector, len(v))
	copy(clone, v)
	return clone
}

// Size returns the number of elements in the Vector.
func (v Vector) Size() int {
	return len(v)
}

// ZerosLike returns a zero-initialized Vector with the same size.
func (v Vector) ZerosLike() Vector {
	return ZeroVector(len(v))
}

// AddScalarInPlace adds a scalar to each element of the Vector in place.
func (v Vector) AddScalarInPlace(scalar float64) Vector {
	for i := range v {
		v[i] += scalar
	}
	return v
}

// AddScalar adds a scalar to each element and returns a new Vector.
func (v Vector) AddScalar(scalar float64) Vector {
	return v.Clone().AddScalarInPlace(scalar)
}

// MulScalarInPlace multiplies each element by a scalar in place.
func (v Vector) MulScalarInPlace(scalar float64) Vector {
	for i := range v {
		v[i] *= scalar
	}
	return v
}

// MulScalar multiplies each element by a scalar and returns a new Vector.
func (v Vector) MulScalar(scalar float64) Vector {
	return v.Clone().MulScalarInPlace(scalar)
}

// Max returns the maximum value in the Vector.
func (v Vector) Max() float64 {
	maxValue := math.Inf(-1)
	for _, value := range v {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// Min returns the minimum value in the Vector.
func (v Vector) Min() float64 {
	minValue := math.Inf(1)
	for _, value := range v {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

// MinMax returns the minimum and maximum values in the Vector.
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
