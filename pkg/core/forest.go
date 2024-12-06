package core

import (
	"math"
	"math/rand/v2"
)

const (
	defaultNumTrees   = 100
	defaultSampleSize = 256
)

type IsolationForest struct {
	Trees []*TreeNode

	// threshold float64

	numTrees    int
	sampleSize  int
	heightLimit int
}

func NewIsolationForest() *IsolationForest {
	f := &IsolationForest{}
	f.Initialize()
	return f
}

func (f *IsolationForest) Initialize() {
	if f.numTrees == 0 {
		f.numTrees = defaultNumTrees
	}

	if f.sampleSize == 0 {
		f.sampleSize = defaultSampleSize
	}

	if f.heightLimit == 0 {
		f.heightLimit = int(math.Ceil(math.Log2(float64(f.sampleSize))))
	}
}

func (forest *IsolationForest) Fit(data [][]float64) {
	for i := 0; i < forest.numTrees; i++ {
		sampledData := sample(data, forest.sampleSize)

		tree := forest.BuildTree(sampledData, 0)
		forest.Trees = append(forest.Trees, tree)
	}
}

func (f *IsolationForest) BuildTree(data [][]float64, currentHeight int) *TreeNode {
	if currentHeight >= f.heightLimit || len(data) <= 1 {
		return &TreeNode{Size: len(data)}
	}

	splitAttribute := rand.IntN(len(data[0]))

	maxValue := maxValue(slice(data, splitAttribute))
	minValue := minValue(slice(data, splitAttribute))

	splitValue := rand.Float64()*(maxValue-minValue) + minValue

	leftData := [][]float64{}
	rightData := [][]float64{}
	for _, row := range data {
		if row[splitAttribute] < splitValue {
			leftData = append(leftData, row)
		} else {
			rightData = append(rightData, row)
		}
	}

	return &TreeNode{
		Left:           f.BuildTree(leftData, currentHeight+1),
		Right:          f.BuildTree(rightData, currentHeight+1),
		SplitAttribute: splitAttribute,
		SplitValue:     splitValue,
	}

}

func (f *IsolationForest) pathLength(vector []float64, node *TreeNode, currentPathLength int) float64 {
	if node.IsLeaf() {
		return float64(currentPathLength) + averagePathLength(node.Size)
	}

	splitAttribute := node.SplitAttribute
	splitValue := node.SplitValue
	if vector[splitAttribute] < splitValue {
		return f.pathLength(vector, node.Left, currentPathLength+1)
	} else {
		return f.pathLength(vector, node.Right, currentPathLength+1)
	}
}

func (f *IsolationForest) Score(data [][]float64) []float64 {
	scores := make([]float64, len(data))

	if len(scores) != len(data) {
		panic("data and scores must have the same length")
	}

	for _, tree := range f.Trees {
		for i, vector := range data {
			scores[i] += f.pathLength(vector, tree, 0)
		}
	}
	// return 2.0 ** (-s / self.average_path_length(self.sample_size))
	return scores
}
