package iforest

import (
	"math/rand"
)

type Matrix []Vector

func ZeroMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = ZeroVector(cols)
	}
	return matrix
}

func RandomMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = RandomVector(cols)
	}
	return matrix
}

func (m Matrix) Clone() Matrix {
	clone := make(Matrix, len(m))
	for i := range m {
		clone[i] = m[i].Clone()
	}
	return clone
}

func (m Matrix) Shape() Shape {
	if len(m) == 0 {
		return Shape{0, 0}
	}
	return Shape{len(m), len(m[0])}
}

func (m Matrix) Size(dim int) int {
	return m.Shape().Size(dim)
}

func (m Matrix) ZerosLike() Matrix {
	shape := m.Shape()
	return ZeroMatrix(shape[0], shape[1])
}

func (m Matrix) AddScalarInPlace(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] += scalar
		}
	}
	return m
}

func (m Matrix) AddScalar(scalar float64) Matrix {
	return m.Clone().AddScalarInPlace(scalar)
}

func (m Matrix) MulScalarInPlace(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] *= scalar
		}
	}
	return m
}

func (m Matrix) MulScalar(scalar float64) Matrix {
	return m.Clone().MulScalarInPlace(scalar)
}

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

func (m Matrix) Column(colIndex int) Vector {
	column := ZeroVector(len(m))
	for i, row := range m {
		column[i] = row[colIndex]
	}
	return column
}
