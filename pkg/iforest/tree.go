package iforest

type TreeNode struct {
	Left       *TreeNode
	Right      *TreeNode
	Size       int
	SplitIndex int
	SplitValue float64
}

func (node *TreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (node *TreeNode) traceSplitIndex(sample Vector, indices []int) []int {
	if node.IsLeaf() {
		return indices
	}

	if sample[node.SplitIndex] < node.SplitValue {
		indices = append(indices, node.SplitIndex)
		return node.Left.traceSplitIndex(sample, indices)
	} else {
		indices = append(indices, node.SplitIndex)
		return node.Right.traceSplitIndex(sample, indices)
	}
}

func (node *TreeNode) FeatureImportance(sample Vector) []int {
	indices := node.traceSplitIndex(sample, []int{})
	importance := make([]int, len(sample))
	for _, index := range indices {
		importance[index]++
	}
	return importance
}
