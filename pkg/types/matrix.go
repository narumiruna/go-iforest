package types

import (
	"math/rand/v2"
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

func (m Matrix) AddScalar(scalar float64) Matrix {
	for i := range m {
		m[i] = m[i].AddScalar(scalar)
	}
	return m
}

func (m Matrix) MulScalar(scalar float64) Matrix {
	for i := range m {
		m[i] = m[i].MulScalar(scalar)
	}
	return m
}

func (m Matrix) Sample(sampleSize int) Matrix {
	if len(m) <= sampleSize {
		return m
	}

	perm := rand.Perm(len(m))
	sampledData := make(Matrix, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sampledData[i] = m[perm[i]]
	}
	return sampledData
}

func (m Matrix) Slice(j int) Vector {
	slicedData := ZeroVector(len(m))
	for i, row := range m {
		slicedData[i] = row[j]
	}
	return slicedData
}
