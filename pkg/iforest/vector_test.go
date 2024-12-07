package iforest

import (
	"testing"
)

// TestZeroVector verifies that ZeroVector returns a Vector of the correct size with all elements initialized to zero.
func TestZeroVector(t *testing.T) {
	v := ZeroVector(5)
	expected := Vector{0, 0, 0, 0, 0}
	for i, val := range v {
		if val != expected[i] {
			t.Errorf("Expected %v, got %v", expected, v)
		}
	}
}

// TestRandomVector checks that RandomVector returns a Vector of the correct size with random elements.
func TestRandomVector(t *testing.T) {
	v := RandomVector(5)
	if len(v) != 5 {
		t.Errorf("Expected length 5, got %d", len(v))
	}
}

// TestSize verifies that the Size method returns the correct length of the Vector.
func TestSize(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Size() != 3 {
		t.Errorf("Expected size 3, got %d", v.Size())
	}
}

// TestAddScalar ensures that AddScalar correctly adds a scalar value to each element of the Vector.
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

// TestMulScalar ensures that MulScalar correctly multiplies each element of the Vector by a scalar value.
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

// TestMax verifies that the Max method returns the largest value in the Vector.
func TestMax(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Max() != 3 {
		t.Errorf("Expected max 3, got %f", v.Max())
	}
}

// TestMin verifies that the Min method returns the smallest value in the Vector.
func TestMin(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Min() != 1 {
		t.Errorf("Expected min 1, got %f", v.Min())
	}
}
