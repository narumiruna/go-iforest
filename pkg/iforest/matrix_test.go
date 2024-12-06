package iforest

import (
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	m := ZeroMatrix(2, 3)
	if len(m) != 2 || len(m[0]) != 3 {
		t.Errorf("Expected shape (2, 3), got (%d, %d)", len(m), len(m[0]))
	}
	for i := range m {
		for j := range m[i] {
			if m[i][j] != 0 {
				t.Errorf("Expected all elements to be 0, got %f at (%d, %d)", m[i][j], i, j)
			}
		}
	}
}

func TestRandomMatrix(t *testing.T) {
	m := RandomMatrix(2, 3)
	if len(m) != 2 || len(m[0]) != 3 {
		t.Errorf("Expected shape (2, 3), got (%d, %d)", len(m), len(m[0]))
	}
}

func TestMatrixShape(t *testing.T) {
	m := ZeroMatrix(2, 3)
	rows, cols := m.Size(0), m.Size(1)
	if rows != 2 || cols != 3 {
		t.Errorf("Expected shape (2, 3), got (%d, %d)", rows, cols)
	}
}

func TestMatrixSize(t *testing.T) {
	m := ZeroMatrix(2, 3)
	if m.Size(0) != 2 || m.Size(1) != 3 {
		t.Errorf("Expected sizes 2 and 3, got %d and %d", m.Size(0), m.Size(1))
	}
}

func TestMatrixAddScalar(t *testing.T) {
	m := ZeroMatrix(2, 3)
	m = m.AddScalar(5)
	for i := range m {
		for j := range m[i] {
			if m[i][j] != 5 {
				t.Errorf("Expected all elements to be 5, got %f", m[i][j])
			}
		}
	}
}

func TestMatrixMulScalar(t *testing.T) {
	m := ZeroMatrix(2, 3)
	m = m.AddScalar(2).MulScalar(3)
	for i := range m {
		for j := range m[i] {
			if m[i][j] != 6 {
				t.Errorf("Expected all elements to be 6, got %f", m[i][j])
			}
		}
	}
}

func TestMatrixSample(t *testing.T) {
	m := RandomMatrix(10, 3)
	sample := m.Sample(5)
	if len(sample) != 5 {
		t.Errorf("Expected sample size 5, got %d", len(sample))
	}
}

func TestMatrixSlice(t *testing.T) {
	m := ZeroMatrix(2, 3)
	m[0][1] = 1
	m[1][1] = 2
	slice := m.Column(1)
	if slice[0] != 1 || slice[1] != 2 {
		t.Errorf("Expected slice [1, 2], got [%f, %f]", slice[0], slice[1])
	}
}
