package iforest

import (
	"math/rand"
)

// Matrix represents a two-dimensional slice of Vectors, effectively a slice of slices of float64 values.
type Matrix []Vector

// ZeroMatrix returns a zero-initialized Matrix with the specified number of rows and columns.
//
// Parameters:
//     rows - the number of rows in the matrix.
//     cols - the number of columns in the matrix.
//
// Returns:
//     A Matrix of size (rows x cols) with all elements initialized to zero.
//
// Example:
//     m := ZeroMatrix(2, 3)  // m == [[0, 0, 0], [0, 0, 0]]
func ZeroMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = ZeroVector(cols)
	}
	return matrix
}

// RandomMatrix returns a Matrix with the specified dimensions containing random values between 0 and 1.
//
// Parameters:
//     rows - the number of rows in the matrix.
//     cols - the number of columns in the matrix.
//
// Returns:
//     A Matrix of size (rows x cols) with random float64 values.
//
// Example:
//     m := RandomMatrix(2, 3)  // m might be [[0.1, 0.2, 0.3], [0.4, 0.5, 0.6]]
func RandomMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = RandomVector(cols)
	}
	return matrix
}

// Clone creates a deep copy of the Matrix.
//
// Returns:
//     A new Matrix that is an exact copy of the original.
//
// Example:
//     m := Matrix{Vector{1, 2}, Vector{3, 4}}
//     clone := m.Clone()  // clone == [[1,2], [3,4]]
func (m Matrix) Clone() Matrix {
	clone := make(Matrix, len(m))
	for i := range m {
		clone[i] = m[i].Clone()
	}
	return clone
}

// Shape returns the dimensions of the Matrix as a Shape.
//
// Returns:
//     A Shape containing two integers: [number of rows, number of columns].
//
// Example:
//     m := ZeroMatrix(2, 3)
//     shape := m.Shape()  // shape == [2, 3]
func (m Matrix) Shape() Shape {
	if len(m) == 0 {
		return Shape{0, 0}
	}
	return Shape{len(m), len(m[0])}
}

// Size returns the size of the specified dimension.
//
// Parameters:
//     dim - the dimension index (0 for rows, 1 for columns).
//
// Returns:
//     An integer representing the size of the specified dimension.
//
// Panics:
//     If 'dim' is not 0 or 1.
//
// Example:
//     m := ZeroMatrix(2, 3)
//     rows := m.Size(0)  // rows == 2
//     cols := m.Size(1)  // cols == 3
func (m Matrix) Size(dim int) int {
	return m.Shape().Size(dim)
}

// ZerosLike returns a zero-initialized Matrix with the same shape as the original.
//
// Returns:
//     A new Matrix with the same dimensions as 'm', with all elements set to zero.
//
// Example:
//     m := RandomMatrix(2, 3)
//     z := m.ZerosLike()  // z == [[0, 0, 0], [0, 0, 0]]
func (m Matrix) ZerosLike() Matrix {
	shape := m.Shape()
	return ZeroMatrix(shape[0], shape[1])
}

// AddScalarInPlace adds a scalar value to each element of the Matrix in place.
//
// Parameters:
//     scalar - the float64 value to add to each element.
//
// Returns:
//     The modified Matrix with updated values.
//
// Example:
//     m := ZeroMatrix(2, 3)
//     m.AddScalarInPlace(1.0)  // m == [[1, 1, 1], [1, 1, 1]]
func (m Matrix) AddScalarInPlace(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] += scalar
		}
	}
	return m
}

// AddScalar adds a scalar to each element of the Matrix and returns a new Matrix.
//
// Parameters:
//     scalar - the float64 value to add to each element.
//
// Returns:
//     A new Matrix with the result of the addition.
//
// Example:
//     m := ZeroMatrix(2, 3)
//     result := m.AddScalar(1.0)  // result == [[1, 1, 1], [1, 1, 1]]
func (m Matrix) AddScalar(scalar float64) Matrix {
	return m.Clone().AddScalarInPlace(scalar)
}

// MulScalarInPlace multiplies each element of the Matrix by a scalar in place.
//
// Parameters:
//     scalar - the float64 value to multiply each element by.
//
// Returns:
//     The modified Matrix with updated values.
//
// Example:
//     m := RandomMatrix(2, 3)
//     m.MulScalarInPlace(2.0)  // Each element in 'm' is multiplied by 2.0
func (m Matrix) MulScalarInPlace(scalar float64) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] *= scalar
		}
	}
	return m
}

// MulScalar multiplies each element of the Matrix by a scalar and returns a new Matrix.
//
// Parameters:
//     scalar - the float64 value to multiply each element by.
//
// Returns:
//     A new Matrix with the result of the multiplication.
//
// Example:
//     m := RandomMatrix(2, 3)
//     result := m.MulScalar(2.0)  // result has elements multiplied by 2.0
func (m Matrix) MulScalar(scalar float64) Matrix {
	return m.Clone().MulScalarInPlace(scalar)
}

// Sample returns a random sample of rows from the Matrix.
//
// Parameters:
//     sampleSize - the number of rows to sample.
//
// Returns:
//     A new Matrix containing 'sampleSize' randomly selected rows from 'm'.
//
// Panics:
//     If 'sampleSize' is less than or equal to zero.
//
// Example:
//     m := RandomMatrix(100, 3)
//     sample := m.Sample(10)  // sample has 10 rows
func (m Matrix) Sample(sampleSize int) Matrix {
	if sampleSize <= 0 {
		panic("sampleSize must be greater than 0")
	}

	if len(m) <= sampleSize {
		return m.Clone()
	}

	perm := rand.Perm(len(m))
	sample := make(Matrix, sampleSize)
	for i := 0; i < sampleSize; i++ {
		sample[i] = m[perm[i]].Clone()
	}
	return sample
}

// Column returns the specified column from the Matrix as a Vector.
//
// Parameters:
//     colIndex - the index of the column to extract (zero-based).
//
// Returns:
//     A Vector containing the elements of the specified column.
//
// Panics:
//     If 'colIndex' is out of range.
//
// Example:
//     m := Matrix{Vector{1,2,3}, Vector{4,5,6}}
//     col := m.Column(1)  // col == [2, 5]
func (m Matrix) Column(colIndex int) Vector {
	column := ZeroVector(len(m))
	for i, row := range m {
		column[i] = row[colIndex]
	}
	return column
}
