package types

import (
	"testing"
)

func TestZeroVector(t *testing.T) {
	v := ZeroVector(5)
	expected := Vector{0, 0, 0, 0, 0}
	for i, val := range v {
		if val != expected[i] {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	}
}

func TestRandomVector(t *testing.T) {
	v := RandomVector(5)
	if len(v) != 5 {
		t.Errorf("Expected length 5, got %d", len(v))
	}
}

func TestShape(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Size() != 3 {
		t.Errorf("Expected shape 3, got %d", v.Size())
	}
}

func TestAddScalar(t *testing.T) {
	v := Vector{1, 2, 3}
	result := v.AddScalar(2)
	expected := Vector{3, 4, 5}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestMulScalar(t *testing.T) {
	v := Vector{1, 2, 3}
	result := v.MulScalar(2)
	expected := Vector{2, 4, 6}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestMax(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Max() != 3 {
		t.Errorf("Expected max 3, got %f", v.Max())
	}
}

func TestMin(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Min() != 1 {
		t.Errorf("Expected min 1, got %f", v.Min())
	}
}
