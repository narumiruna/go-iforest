package iforest

// Shape represents the dimensions of a matrix or multi-dimensional array.
// It is a slice of integers where each element specifies the size of the corresponding dimension.
type Shape []int

// Dim returns the number of dimensions represented by the Shape.
// Returns:
//     An integer indicating the number of dimensions.
// Example:
//     s := Shape{2, 3}
//     n := s.Dim()  // n == 2
func (s Shape) Dim() int {
	return len(s)
}

// Size returns the size of the specified dimension.
// Parameters:
//     i - the index of the dimension (zero-based).
// Returns:
//     An integer representing the size of the i-th dimension.
// Panics:
//     If i is out of range.
// Example:
//     s := Shape{2, 3}
//     size := s.Size(1)  // size == 3
func (s Shape) Size(i int) int {
	if i < 0 || i >= len(s) {
		panic("index out of range")
	}
	return s[i]
}
