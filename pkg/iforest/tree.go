package iforest

// TreeNode represents a node in the isolation tree.
type TreeNode struct {
	Left       *TreeNode
	Right      *TreeNode
	Size       int
	SplitIndex int
	SplitValue float64
}

// IsLeaf returns true if the node is a leaf.
func (t *TreeNode) IsLeaf() bool {
	return t.Left == nil && t.Right == nil
}

// traceSplitIndices traces the split indices for a sample.
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

// FeatureImportance computes feature importance for a sample.
func (t *TreeNode) FeatureImportance(sample Vector) []int {
	indices := t.traceSplitIndices(sample, []int{})
	importance := make([]int, len(sample))
	for _, index := range indices {
		importance[index]++
	}
	return importance
}
