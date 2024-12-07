// Package iforest provides implementations of isolation forest algorithms for anomaly detection.
//
// This package includes data structures and utility functions for handling vectors and matrices,
// which are fundamental components used in building the isolation forest model.
package iforest

import (
	"math"
	"math/rand"
)

// Vector represents a one-dimensional array of float64 values, commonly used to store feature vectors.
type Vector []float64

// ZeroVector returns a zero-initialized Vector of the specified size.
//
// Parameters:
//
//	size - the length of the vector to create.
//
// Returns:
//
//	A Vector of length 'size' with all elements initialized to zero.
//
// Example:
//
//	v := ZeroVector(3)  // v == [0.0, 0.0, 0.0]
func ZeroVector(size int) Vector {
	return make(Vector, size)
}

// RandomVector returns a Vector of the specified size with random float64 values between 0 and 1.
//
// Parameters:
//
//	size - the length of the vector to create.
//
// Returns:
//
//	A Vector of length 'size' with random values.
//
// Example:
//
//	v := RandomVector(3)  // v might be [0.123, 0.456, 0.789]
func RandomVector(size int) Vector {
	vector := make(Vector, size)
	for i := range vector {
		vector[i] = rand.Float64()
	}
	return vector
}

// Clone creates a deep copy of the Vector.
//
// Returns:
//
//	A new Vector that is a copy of the original.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	clone := v.Clone()  // clone == [1.0, 2.0, 3.0]
func (v Vector) Clone() Vector {
	clone := make(Vector, len(v))
	copy(clone, v)
	return clone
}

// Size returns the number of elements in the Vector.
//
// Returns:
//
//	An integer representing the length of the vector.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	n := v.Size()  // n == 3
func (v Vector) Size() int {
	return len(v)
}

// ZerosLike returns a zero-initialized Vector with the same size as the original Vector.
//
// Returns:
//
//	A new Vector of the same length with all elements set to zero.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	z := v.ZerosLike()  // z == [0.0, 0.0, 0.0]
func (v Vector) ZerosLike() Vector {
	return ZeroVector(len(v))
}

// AddScalarInPlace adds a scalar value to each element of the Vector in place.
//
// Parameters:
//
//	scalar - the float64 value to add to each element.
//
// Returns:
//
//	The modified Vector with updated values.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	v.AddScalarInPlace(2.0)  // v == [3.0, 4.0, 5.0]
func (v Vector) AddScalarInPlace(scalar float64) Vector {
	for i := range v {
		v[i] += scalar
	}
	return v
}

// AddScalar adds a scalar to each element of the Vector and returns a new Vector.
//
// Parameters:
//
//	scalar - the float64 value to add to each element.
//
// Returns:
//
//	A new Vector with the result of the addition.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	result := v.AddScalar(2.0)  // result == [3.0, 4.0, 5.0]
func (v Vector) AddScalar(scalar float64) Vector {
	return v.Clone().AddScalarInPlace(scalar)
}

// MulScalarInPlace multiplies each element of the Vector by a scalar in place.
//
// Parameters:
//
//	scalar - the float64 value to multiply each element by.
//
// Returns:
//
//	The modified Vector with updated values.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	v.MulScalarInPlace(2.0)  // v == [2.0, 4.0, 6.0]
func (v Vector) MulScalarInPlace(scalar float64) Vector {
	for i := range v {
		v[i] *= scalar
	}
	return v
}

// MulScalar multiplies each element of the Vector by a scalar and returns a new Vector.
//
// Parameters:
//
//	scalar - the float64 value to multiply each element by.
//
// Returns:
//
//	A new Vector with the result of the multiplication.
//
// Example:
//
//	v := Vector{1.0, 2.0, 3.0}
//	result := v.MulScalar(2.0)  // result == [2.0, 4.0, 6.0]
func (v Vector) MulScalar(scalar float64) Vector {
	return v.Clone().MulScalarInPlace(scalar)
}

// Max returns the maximum value in the Vector.
//
// Returns:
//
//	The largest float64 value found in the Vector. If the Vector is empty, returns math.Inf(-1).
//
// Example:
//
//	v := Vector{1.0, 5.0, 3.0}
//	maxVal := v.Max()  // maxVal == 5.0
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
//
// Returns:
//
//	The smallest float64 value found in the Vector. If the Vector is empty, returns math.Inf(1).
//
// Example:
//
//	v := Vector{1.0, 5.0, 3.0}
//	minVal := v.Min()  // minVal == 1.0
func (v Vector) Min() float64 {
	minValue := math.Inf(1)
	for _, value := range v {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

// MinMax returns both the minimum and maximum values in the Vector.
//
// Returns:
//
//	minValue - the smallest float64 value in the Vector.
//	maxValue - the largest float64 value in the Vector.
//	If the Vector is empty, returns (math.Inf(1), math.Inf(-1)).
//
// Example:
//
//	v := Vector{1.0, 5.0, 3.0}
//	minVal, maxVal := v.MinMax()  // minVal == 1.0, maxVal == 5.0
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
