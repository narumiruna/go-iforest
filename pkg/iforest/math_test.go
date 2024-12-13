package iforest

import (
	"testing"
)

func TestSample(t *testing.T) {
	matrix := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
		{5.0, 6.0},
		{7.0, 8.0},
	}
	sampleSize := 2
	sample := Sample(matrix, sampleSize)
	if len(sample) != sampleSize {
		t.Errorf("Expected sample size %d, got %d", sampleSize, len(sample))
	}
}

func TestColumn(t *testing.T) {
	matrix := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
		{5.0, 6.0},
	}
	column := Column(matrix, 1)
	expected := []float64{2.0, 4.0, 6.0}
	for i, v := range column {
		if v != expected[i] {
			t.Errorf("Expected %f, got %f", expected[i], v)
		}
	}
}

func TestMinMax(t *testing.T) {
	slice := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	min, max := MinMax(slice)
	if min != 1.0 {
		t.Errorf("Expected min 1.0, got %f", min)
	}
	if max != 5.0 {
		t.Errorf("Expected max 5.0, got %f", max)
	}
}

func TestRandomMatrix(t *testing.T) {
	rows, cols := 3, 2
	matrix := RandomMatrix(rows, cols)
	if len(matrix) != rows {
		t.Errorf("Expected %d rows, got %d", rows, len(matrix))
	}
	for _, row := range matrix {
		if len(row) != cols {
			t.Errorf("Expected %d columns, got %d", cols, len(row))
		}
	}
}

func TestAddScalar(t *testing.T) {
	matrix := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	scalar := 1.5
	result := AddScalar(matrix, scalar)
	expected := [][]float64{
		{2.5, 3.5},
		{4.5, 5.5},
	}
	for i, row := range result {
		for j, v := range row {
			if v != expected[i][j] {
				t.Errorf("Expected %f, got %f", expected[i][j], v)
			}
		}
	}
}
