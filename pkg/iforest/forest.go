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
	ScoreThreshold float64
	NumTrees       int
	SampleSize     int
	MaxDepth       int

	Trees []*TreeNode
}

type IsolationForestOption struct {
	ScoreThreshold float64 `json:"score_threshold"`
	NumTrees       int     `json:"num_trees"`
	SampleSize     int     `json:"sample_size"`
	MaxDepth       int     `json:"max_depth"`
}

func New() *IsolationForest {
	f := &IsolationForest{}
	f.setDefaultValues()
	return f
}

func NewWithOptions(options IsolationForestOption) *IsolationForest {
	f := &IsolationForest{
		ScoreThreshold: options.ScoreThreshold,
		NumTrees:       options.NumTrees,
		SampleSize:     options.SampleSize,
		MaxDepth:       options.MaxDepth,
	}
	f.setDefaultValues()
	return f
}

func (f *IsolationForest) setDefaultValues() {
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

	leftSamples := Matrix{}
	rightSamples := Matrix{}
	for _, vector := range samples {
		if vector[splitIndex] < splitValue {
			leftSamples = append(leftSamples, vector)
		} else {
			rightSamples = append(rightSamples, vector)
		}
	}

	return &TreeNode{
		Left:       f.BuildTree(leftSamples, depth+1),
		Right:      f.BuildTree(rightSamples, depth+1),
		SplitIndex: splitIndex,
		SplitValue: splitValue,
	}

}

func (f *IsolationForest) Score(samples Matrix) []float64 {
	scores := make([]float64, len(samples))
	for i, sample := range samples {
		s := 0.0
		for _, tree := range f.Trees {
			s += pathLength(sample, tree, 0)
		}
		scores[i] = math.Pow(2.0, -s/float64(len(f.Trees))/averagePathLength(float64(f.SampleSize)))
	}
	return scores
}

func (f *IsolationForest) Predict(samples Matrix) []int {
	predicts := make([]int, len(samples))
	for i, s := range f.Score(samples) {
		if s > f.ScoreThreshold {
			predicts[i] = 1
		} else {
			predicts[i] = 0
		}

	}
	return predicts
}

func (f *IsolationForest) FeatureImportance(sample Vector) []int {
	importance := make([]int, len(sample))
	for _, tree := range f.Trees {
		for i, v := range tree.FeatureImportance(sample) {
			importance[i] += v
		}
	}
	return importance
}
