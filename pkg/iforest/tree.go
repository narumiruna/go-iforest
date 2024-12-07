package iforest

// TreeNode represents a node in the isolation tree used by the isolation forest algorithm.
//
// Each TreeNode contains information about the feature to split on, the split value, and references to its child nodes.
type TreeNode struct {
	Left       *TreeNode
	Right      *TreeNode
	Size       int
	SplitIndex int
	SplitValue float64
}

// IsLeaf returns true if the node is a leaf node (i.e., has no children).
//
// Returns:
//     A boolean indicating whether the node is a leaf.
//
// Example:
//     node := &TreeNode{}
//     isLeaf := node.IsLeaf()  // isLeaf == true
func (t *TreeNode) IsLeaf() bool {
	return t.Left == nil && t.Right == nil
}

// traceSplitIndices traverses the tree to collect the indices of features used for splitting for a given sample.
//
// Parameters:
//     sample - the input Vector representing a data point.
//     indices - slice to accumulate the feature indices used during traversal.
//
// Returns:
//     A slice of integers representing the feature indices used during traversal.
func (t *TreeNode) traceSplitIndices(sample Vector, indices []int) []int {
	if t.IsLeaf() {
		return indices
	}

	if sample[t.SplitIndex] < t.SplitValue {
		indices = append(indices, t.SplitIndex)
		return t.Left.traceSplitIndices(sample, indices)
	} else {
		indices = append(indices, t.SplitIndex)
		return t.Right.traceSplitIndices(sample, indices)
	}
}

// FeatureImportance computes the importance of features for a given sample based on the path taken in the tree.
//
// Parameters:
//     sample - the input Vector representing a data point.
//
// Returns:
//     A slice of integers where each element represents the importance (frequency) of the corresponding feature.
func (t *TreeNode) FeatureImportance(sample Vector) []int {
	indices := t.traceSplitIndices(sample, []int{})
	importance := make([]int, len(sample))
	for _, index := range indices {
		importance[index]++
	}
	return importance
}
