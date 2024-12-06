package iforest

type TreeNode struct {
	Left           *TreeNode
	Right          *TreeNode
	Size           int
	SplitAttribute int
	SplitValue     float64
}

func (node *TreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}
