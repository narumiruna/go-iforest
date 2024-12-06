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
	Trees          []*TreeNode
	ScoreThreshold float64
	NumTrees       int
	SampleSize     int
	HeightLimit    int
}

func NewIsolationForest() *IsolationForest {
	f := &IsolationForest{}
	f.Initialize()
	return f
}

func (f *IsolationForest) Initialize() {
	if f.ScoreThreshold == 0 {
		f.ScoreThreshold = defaultScoreThreshold
	}

	if f.NumTrees == 0 {
		f.NumTrees = defaultNumTrees
	}

	if f.SampleSize == 0 {
		f.SampleSize = defaultSampleSize
	}

	if f.HeightLimit == 0 {
		f.HeightLimit = int(math.Ceil(math.Log2(float64(f.SampleSize))))
	}
}

func (f *IsolationForest) Fit(data Matrix) {
	for i := 0; i < f.NumTrees; i++ {
		sampledData := data.Sample(f.SampleSize)
		tree := f.BuildTree(sampledData, 0)
		f.Trees = append(f.Trees, tree)
	}
}

func (f *IsolationForest) BuildTree(data Matrix, currentHeight int) *TreeNode {
	numSamples, numFeatures := data.Size(0), data.Size(1)
	if currentHeight >= f.HeightLimit || numSamples <= 1 {
		return &TreeNode{Size: numSamples}
	}

	splitAttribute := rand.Intn(numFeatures)
	slicedData := data.Column(splitAttribute)
	maxValue := slicedData.Max()
	minValue := slicedData.Min()

	splitValue := rand.Float64()*(maxValue-minValue) + minValue

	leftData := Matrix{}
	rightData := Matrix{}
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
