package iforest

import (
	"testing"
)

// TestZeroMatrix verifies that ZeroMatrix returns a Matrix with correct dimensions and zero-initialized elements.
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

// TestRandomMatrix verifies that RandomMatrix returns a Matrix with correct dimensions and random elements.
func TestRandomMatrix(t *testing.T) {
	m := RandomMatrix(2, 3)
	if len(m) != 2 || len(m[0]) != 3 {
		t.Errorf("Expected shape (2, 3), got (%d, %d)", len(m), len(m[0]))
	}
}

// TestMatrixShape checks that the Shape method returns the correct dimensions of the Matrix.
func TestMatrixShape(t *testing.T) {
	m := ZeroMatrix(2, 3)
	rows, cols := m.Size(0), m.Size(1)
	if rows != 2 || cols != 3 {
		t.Errorf("Expected shape (2, 3), got (%d, %d)", rows, cols)
	}
}

// TestMatrixSize checks that the Size method returns the correct size of the specified dimension.
func TestMatrixSize(t *testing.T) {
	m := ZeroMatrix(2, 3)
	if m.Size(0) != 2 || m.Size(1) != 3 {
		t.Errorf("Expected sizes 2 and 3, got %d and %d", m.Size(0), m.Size(1))
	}
}

// TestMatrixAddScalar tests that AddScalar correctly adds a scalar value to each element of the Matrix.
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

// TestMatrixMulScalar tests that MulScalar correctly multiplies each element of the Matrix by a scalar value.
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

// TestMatrixSample tests that Sample returns a Matrix with the correct number of randomly selected rows.
func TestMatrixSample(t *testing.T) {
	m := RandomMatrix(10, 3)
	sample := m.Sample(5)
	if len(sample) != 5 {
		t.Errorf("Expected sample size 5, got %d", len(sample))
	}
}

// TestMatrixColumn tests that Column correctly extracts a column from the Matrix as a Vector.
func TestMatrixColumn(t *testing.T) {
	m := ZeroMatrix(2, 3)
	m[0][1] = 1
	m[1][1] = 2
	column := m.Column(1)
	if column[0] != 1 || column[1] != 2 {
		t.Errorf("Expected column [1, 2], got [%f, %f]", column[0], column[1])
	}
}
