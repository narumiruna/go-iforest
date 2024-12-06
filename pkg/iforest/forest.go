package iforest

import (
	"math"
	"math/rand"
)

const (
	defaultNumTrees       = 100
	defaultSampleSize     = 256
	defaultScoreThreshold = 0.6
)

type IsolationForest struct {
	ScoreThreshold float64 `json:"score_threshold"`
	NumTrees       int     `json:"num_trees"`
	SampleSize     int     `json:"sample_size"`
	MaxDepth       int     `json:"max_depth"`

	Trees []*TreeNode
}

func New() *IsolationForest {
	f := &IsolationForest{}
	f.init()
	return f
}

func (f *IsolationForest) init() {
	if f.ScoreThreshold == 0 {
		f.ScoreThreshold = defaultScoreThreshold
	}

	if f.NumTrees == 0 {
		f.NumTrees = defaultNumTrees
	}

	if f.SampleSize == 0 {
		f.SampleSize = defaultSampleSize
	}

	if f.MaxDepth == 0 {
		f.MaxDepth = int(math.Ceil(math.Log2(float64(f.SampleSize))))
	}
}

func (f *IsolationForest) Fit(samples Matrix) {
	for i := 0; i < f.NumTrees; i++ {
		sampled := samples.Sample(f.SampleSize)
		tree := f.BuildTree(sampled, 0)
		f.Trees = append(f.Trees, tree)
	}
}

func (f *IsolationForest) BuildTree(samples Matrix, depth int) *TreeNode {
	numSamples, numFeatures := samples.Size(0), samples.Size(1)
	if depth >= f.MaxDepth || numSamples <= 1 {
		return &TreeNode{Size: numSamples}
	}

	splitIndex := rand.Intn(numFeatures)
	column := samples.Column(splitIndex)
	maxValue := column.Max()
	minValue := column.Min()
	splitValue := rand.Float64()*(maxValue-minValue) + minValue

	leftData := Matrix{}
	rightData := Matrix{}
	for _, vector := range samples {
		if vector[splitIndex] < splitValue {
			leftData = append(leftData, vector)
		} else {
			rightData = append(rightData, vector)
		}
	}

	return &TreeNode{
		Left:       f.BuildTree(leftData, depth+1),
		Right:      f.BuildTree(rightData, depth+1),
		SplitIndex: splitIndex,
		SplitValue: splitValue,
	}

}

func (f *IsolationForest) Score(data Matrix) []float64 {
	scores := ZeroVector(len(data))

	for i, vector := range data {
		s := 0.0
		for _, tree := range f.Trees {
			s += pathLength(vector, tree, 0)
		}
		scores[i] = math.Pow(2.0, -s/float64(len(f.Trees))/averagePathLength(float64(f.SampleSize)))
	}

	return scores
}

func (f *IsolationForest) Predict(data Matrix) []int {
	predicts := make([]int, len(data))

	scores := f.Score(data)

	ScoreThreshold := f.ScoreThreshold
	for i, s := range scores {
		if s > ScoreThreshold {
			predicts[i] = 1
		} else {
			predicts[i] = 0
		}

	}
	return predicts
}
