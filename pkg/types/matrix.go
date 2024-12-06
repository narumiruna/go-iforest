package types

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

func (m Matrix) Shape() (int, int) {
	if len(m) == 0 {
		return 0, 0
	}
	return len(m), len(m[0])
}

func (m Matrix) Size(i int) int {
	s1, s2 := m.Shape()
	if i == 0 {
		return s1
	}
	return s2
}

func (m Matrix) ZerosLike() Matrix {
	p, q := m.Shape()
	return ZeroMatrix(p, q)
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
