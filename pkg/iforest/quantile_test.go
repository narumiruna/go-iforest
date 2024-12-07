package iforest

import (
	"testing"
)

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
		{[]float64{1, 2, 3, 4, 5}, 0.001, 1},
	}

	for _, c := range cases {
		actual := Quantile(c.numbers, c.q)
		if actual != c.expected {
			t.Errorf("Quantile(%v, %v) == %v, expected %v", c.numbers, c.q, actual, c.expected)
		}
	}
}
