package iforest

import "math"

const EulerGamma = 0.5772156649

// harmonicNumber computes the harmonic number of x.
func harmonicNumber(x float64) float64 {
	return math.Log(x) + EulerGamma
}

// averagePathLength computes the average path length for given x.
func averagePathLength(x float64) float64 {
	if x > 2 {
		return 2.0*harmonicNumber(x-1) - 2.0*(x-1)/x
	} else if x == 2 {
		return 1.0
	} else {
		return 0.0
	}
}

// pathLength calculates the path length of a sample in the tree.
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
