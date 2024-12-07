package iforest

import (
	"math/rand"
)

type Matrix []Vector

func ZeroMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := range matrix {
		matrix[i] = ZeroVector(n)
	}
	return matrix
}

func RandomMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := range matrix {
		matrix[i] = RandomVector(n)
	}
	return matrix
}

func (m Matrix) Clone() Matrix {
	o := make(Matrix, len(m))
	for i := range m {
		o[i] = m[i].Clone()
	}
	return o
}

func (m Matrix) Shape() Shape {
	if len(m) == 0 {
		return Shape{0, 0}
	}
	return Shape{len(m), len(m[0])}
}

func (m Matrix) Size(i int) int {
	return m.Shape().Size(i)
}

func (m Matrix) ZerosLike() Matrix {
	shape := m.Shape()
	return ZeroMatrix(shape[0], shape[1])
}

func (m Matrix) AddScalar_(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] += scalar
		}
	}
	return m
}

func (m Matrix) AddScalar(scalar float64) Matrix {
	return m.Clone().AddScalar_(scalar)
}

func (m Matrix) MulScalar_(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] *= scalar
		}
	}
	return m
}

func (m Matrix) MulScalar(scalar float64) Matrix {
	return m.Clone().MulScalar_(scalar)
}

func (m Matrix) Sample(sampleSize int) Matrix {
	if sampleSize <= 0 {
		panic("sampleSize must be greater than 0")
	}

	if len(m) <= sampleSize {
		return m.Clone()
	}

	perm := rand.Perm(len(m))
	o := make(Matrix, sampleSize)
	for i := 0; i < sampleSize; i++ {
		o[i] = m[perm[i]].Clone()
	}
	return o
}

func (m Matrix) Column(j int) Vector {
	o := ZeroVector(len(m))
	for i, row := range m {
		o[i] = row[j]
	}
	return o
}
