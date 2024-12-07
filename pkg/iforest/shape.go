package iforest

// Shape represents the dimensions of a matrix.
type Shape []int

// Dim returns the number of dimensions.
func (s Shape) Dim() int {
	return len(s)
}

// Size returns the size of the i-th dimension.
func (s Shape) Size(i int) int {
	if i < 0 || i >= len(s) {
		panic("index out of range")
	}
	return s[i]
}
