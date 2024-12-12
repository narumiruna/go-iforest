package iforest

import (
	"fmt"
	"math"
	"sort"
)

// Quantile computes the q-th quantile of a slice of numbers using linear interpolation.
// Parameters:
//     numbers - a slice of float64 values to compute the quantile from.
//     q       - a float64 value between 0 and 1 representing the desired quantile.
// Returns:
//     The computed quantile as a float64.
// Panics:
//     If 'numbers' is empty.
//     If 'q' is not between 0 and 1 (inclusive).
// Example:
//     data := []float64{1, 2, 3, 4, 5}
//     median := Quantile(data, 0.5)  // median == 3.0
func Quantile(numbers []float64, q float64) float64 {
	if len(numbers) == 0 {
		panic("numbers must not be empty")
	}
	if q < 0 || q > 1 {
		panic(fmt.Sprintf("q must be in [0, 1], got %v", q))
	}

	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)

	n := float64(len(sortedNumbers))
	pos := q * (n - 1)
	lowerIndex := int(math.Floor(pos))
	upperIndex := int(math.Ceil(pos))
	if lowerIndex == upperIndex {
		return sortedNumbers[lowerIndex]
	}

	// linear interpolation
	fraction := pos - float64(lowerIndex)
	return sortedNumbers[lowerIndex] + fraction*(sortedNumbers[upperIndex]-sortedNumbers[lowerIndex])
}
