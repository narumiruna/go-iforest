package iforest

import "math"

const EulerGamma = 0.5772156649

// harmonicNumber computes the harmonic number H(x), approximated for real numbers using the natural logarithm and Euler-Mascheroni constant.
//
// Parameters:
//     x - the float64 value to compute the harmonic number for.
//
// Returns:
//     The approximate harmonic number H(x) as a float64, where H(x) ≈ ln(x) + γ.
//
// Example:
//     h := harmonicNumber(5)  // h ≈ 2.283
func harmonicNumber(x float64) float64 {
	return math.Log(x) + EulerGamma
}

// averagePathLength computes the average path length of unsuccessful searches in a binary search tree.
//
// Parameters:
//     x - the float64 value representing the number of external nodes.
//
// Returns:
//     The average path length as a float64.
//
// Example:
//     l := averagePathLength(256)  // l ≈ 9.0
func averagePathLength(x float64) float64 {
	if x > 2 {
		return 2.0*harmonicNumber(x-1) - 2.0*(x-1)/x
	} else if x == 2 {
		return 1.0
	} else {
		return 0.0
	}
}

// pathLength calculates the expected path length of a sample in the isolation tree.
//
// Parameters:
//     vector            - the input Vector representing the sample.
//     node              - the current TreeNode during tree traversal.
//     currentPathLength - the current accumulated path length.
//
// Returns:
//     The path length for the sample as a float64.
//
// This function recursively traverses the tree to compute the path length for the given sample.
func pathLength(vector []float64, node *TreeNode, currentPathLength int) float64 {
	if node.IsLeaf() {
		return float64(currentPathLength) + averagePathLength(float64(node.Size))
	}

	splitAttribute := node.SplitIndex
	splitValue := node.SplitValue
	if vector[splitAttribute] < splitValue {
		return pathLength(vector, node.Left, currentPathLength+1)
	} else {
		return pathLength(vector, node.Right, currentPathLength+1)
	}
}
