package iforest

import (
	"testing"
)

// TestQuantile tests the Quantile function with various datasets and quantile values.
// It iterates over a set of test cases, each containing a slice of numbers, a quantile value, and the expected result.
// The function calculates the quantile for each test case and compares it to the expected result.
// If the calculated quantile does not match the expected result, the test fails and an error message is printed.
func TestQuantile(t *testing.T) {
	cases := []struct {
		numbers  []float64
		q        float64
		expected float64
	}{
		{[]float64{1, 2}, 0.5, 1.5},
		{[]float64{1, 2, 3, 4, 5}, 0.5, 3},
		{[]float64{1, 2, 3, 4, 5}, 1.0, 5},
		{[]float64{1, 2, 3, 4, 5}, 0.0, 1},
	}

	for _, c := range cases {
		actual := Quantile(c.numbers, c.q)
		if actual != c.expected {
			t.Errorf("Quantile(%v, %v) == %v, expected %v", c.numbers, c.q, actual, c.expected)
		}
	}
}
