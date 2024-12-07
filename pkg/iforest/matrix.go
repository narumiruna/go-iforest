package iforest

import (
	"math/rand"
)

// Matrix represents a two-dimensional slice of Vectors.
type Matrix []Vector

// ZeroMatrix returns a zero-initialized Matrix with the given dimensions.
func ZeroMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = ZeroVector(cols)
	}
	return matrix
}

// RandomMatrix returns a Matrix with random values.
func RandomMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = RandomVector(cols)
	}
	return matrix
}

// Clone creates a copy of the Matrix.
func (m Matrix) Clone() Matrix {
	clone := make(Matrix, len(m))
	for i := range m {
		clone[i] = m[i].Clone()
	}
	return clone
}

// Shape returns the dimensions of the Matrix.
func (m Matrix) Shape() Shape {
	if len(m) == 0 {
		return Shape{0, 0}
	}
	return Shape{len(m), len(m[0])}
}

// Size returns the size of the specified dimension.
func (m Matrix) Size(dim int) int {
	return m.Shape().Size(dim)
}

// ZerosLike returns a zero-initialized Matrix with the same shape.
func (m Matrix) ZerosLike() Matrix {
	shape := m.Shape()
	return ZeroMatrix(shape[0], shape[1])
}

// AddScalarInPlace adds a scalar to each element in place.
func (m Matrix) AddScalarInPlace(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] += scalar
		}
	}
	return m
}

// AddScalar adds a scalar to each element and returns a new Matrix.
func (m Matrix) AddScalar(scalar float64) Matrix {
	return m.Clone().AddScalarInPlace(scalar)
}

// MulScalarInPlace multiplies each element by a scalar in place.
func (m Matrix) MulScalarInPlace(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] *= scalar
		}
	}
	return m
}

// MulScalar multiplies each element by a scalar and returns a new Matrix.
func (m Matrix) MulScalar(scalar float64) Matrix {
	return m.Clone().MulScalarInPlace(scalar)
}

// Sample returns a random sample of rows from the Matrix.
func (m Matrix) Sample(sampleSize int) Matrix {
	if sampleSize <= 0 {
		panic("sampleSize must be greater than 0")
	}

	if len(m) <= sampleSize {
		return m.Clone()
	}

	perm := rand.Perm(len(m))
	sample := make(Matrix, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = m[perm[i]].Clone()
	}
	return sample
}

// Column returns the specified column as a Vector.
func (m Matrix) Column(colIndex int) Vector {
	column := ZeroVector(len(m))
	for i, row := range m {
		column[i] = row[colIndex]
	}
	return column
}
