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

func (m Matrix) AddScalar(scalar float64) Matrix {
	o := make(Matrix, len(m))
	for i := range m {
		o[i] = m[i].AddScalar(scalar)
	}
	return o
}

func (m Matrix) MulScalar(scalar float64) Matrix {
	o := make(Matrix, len(m))
	for i := range m {
		o[i] = m[i].MulScalar(scalar)
	}
	return o
}

func (m Matrix) Sample(sampleSize int) Matrix {
	if len(m) <= sampleSize {
		return m
	}

	perm := rand.Perm(len(m))
	o := make(Matrix, sampleSize)
	for i := 0; i < sampleSize; i++ {
		o[i] = m[perm[i]]
	}
	return o
}

func (m Matrix) Column(j int) Vector {
	slicedData := ZeroVector(len(m))
	for i, row := range m {
		slicedData[i] = row[j]
	}
	return slicedData
}
