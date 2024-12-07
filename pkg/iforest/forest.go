package iforest

import (
	"math"
	"math/rand"
	"sync"
)

const (
	defaultNumTrees       = 100
	defaultSampleSize     = 256
	defaultScoreThreshold = 0.6
	defaultDetectionType  = DetectionTypeThreshold
	offset                = 0.5
)

type DetectionType string

const (
	DetectionTypeThreshold  DetectionType = "threshold"
	DetectionTypeProportion DetectionType = "proportion"
)

type IsolationForest struct {
	DetectionType DetectionType `json:"detection_type"`
	Threshold     float64       `json:"threshold"`
	Proportion    float64       `json:"proportion"`
	NumTrees      int           `json:"num_trees"`
	SampleSize    int           `json:"sample_size"`
	MaxDepth      int           `json:"max_depth"`

	Trees []*TreeNode
}

type IsolationForestOption struct {
	DetectionType DetectionType `json:"detection_type"`
	Threshold     float64       `json:"threshold"`
	Proportion    float64       `json:"proportion"`
	NumTrees      int           `json:"num_trees"`
	SampleSize    int           `json:"sample_size"`
	MaxDepth      int           `json:"max_depth"`
}

func New() *IsolationForest {
	f := &IsolationForest{}
	f.setDefaultValues()
	return f
}

func NewWithOptions(options IsolationForestOption) *IsolationForest {
	f := &IsolationForest{
		DetectionType: options.DetectionType,
		Threshold:     options.Threshold,
		Proportion:    options.Proportion,
		NumTrees:      options.NumTrees,
		SampleSize:    options.SampleSize,
		MaxDepth:      options.MaxDepth,
	}
	f.setDefaultValues()
	return f
}

func (f *IsolationForest) setDefaultValues() {
	if f.DetectionType == "" {
		f.DetectionType = defaultDetectionType
	}

	if f.Threshold == 0 {
		f.Threshold = defaultScoreThreshold
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
	wg := sync.WaitGroup{}
	wg.Add(f.NumTrees)

	f.Trees = make([]*TreeNode, f.NumTrees)
	for i := 0; i < f.NumTrees; i++ {
		sampled := samples.Sample(f.SampleSize)
		go func() {
			defer wg.Done()
			tree := f.BuildTree(sampled, 0)
			f.Trees[i] = tree
		}()
	}
	wg.Wait()

}

func (f *IsolationForest) BuildTree(samples Matrix, depth int) *TreeNode {
	numSamples, numFeatures := samples.Size(0), samples.Size(1)
	if depth >= f.MaxDepth || numSamples <= 1 {
		return &TreeNode{Size: numSamples}
	}

	splitIndex := rand.Intn(numFeatures)
	column := samples.Column(splitIndex)
	minValue, maxValue := column.MinMax()
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
		score := 0.0
		for _, tree := range f.Trees {
			score += pathLength(sample, tree, 0)
		}
		scores[i] = math.Pow(2.0, -score/float64(len(f.Trees))/averagePathLength(float64(f.SampleSize)))
	}
	return scores
}

func (f *IsolationForest) Predict(samples Matrix) []int {
	predictions := make([]int, len(samples))
	scores := f.Score(samples)

	var threshold float64
	switch f.DetectionType {
	case DetectionTypeThreshold:
		threshold = f.Threshold
	case DetectionTypeProportion:
		threshold = Quantile(f.Score(samples), 1-f.Proportion)
	default:
		panic("Invalid detection type")

	}

	for i, score := range scores {
		if score > threshold {
			predictions[i] = 1
		} else {
			predictions[i] = 0
		}
	}

	return predictions
}

func (f *IsolationForest) FeatureImportance(sample Vector) []int {
	importance := make([]int, len(sample))
	for _, tree := range f.Trees {
		for i, value := range tree.FeatureImportance(sample) {
			importance[i] += value
		}
	}
	return importance
}
