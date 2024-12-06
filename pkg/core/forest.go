package core

import (
	"math"
	"math/rand/v2"

	"github.com/narumiruna/go-isolation-forest/pkg/types"
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

func (forest *IsolationForest) Fit(data types.Matrix) {
	for i := 0; i < forest.numTrees; i++ {
		sampledData := data.Sample(forest.sampleSize)
		tree := forest.BuildTree(sampledData, 0)
		forest.Trees = append(forest.Trees, tree)
	}
}

func (f *IsolationForest) BuildTree(data types.Matrix, currentHeight int) *TreeNode {
	nRows, nCols := data.Shape()
	if currentHeight >= f.heightLimit || nRows <= 1 {
		return &TreeNode{Size: nRows}
	}

	splitAttribute := rand.IntN(nCols)
	slicedData := data.Slice(splitAttribute)
	maxValue := slicedData.Max()
	minValue := slicedData.Min()

	splitValue := rand.Float64()*(maxValue-minValue) + minValue

	leftData := types.Matrix{}
	rightData := types.Matrix{}
	for _, vector := range data {
		if vector[splitAttribute] < splitValue {
			leftData = append(leftData, vector)
		} else {
			rightData = append(rightData, vector)
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

func (f *IsolationForest) Score(data types.Matrix) []float64 {
	scores := make([]float64, len(data))

	if len(scores) != len(data) {
		panic("data and scores must have the same length")
	}

	for _, tree := range f.Trees {
		for i, vector := range data {
			scores[i] += f.pathLength(vector, tree, 0)
		}
	}

	// average
	for i := range scores {
		scores[i] /= float64(len(f.Trees))
	}

	for i, s := range scores {
		scores[i] = math.Pow(2.0, -s/averagePathLength(len(data)))
	}
	return scores
}

func (f *IsolationForest) Predict(data types.Matrix) []int {
	predicts := make([]int, len(data))

	scores := f.Score(data)

	threshold := 0.6
	for i, s := range scores {
		if s > threshold {
			predicts[i] = 1
		} else {
			predicts[i] = 0
		}

	}
	return predicts
}
