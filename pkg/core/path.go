package core

import "math"

const eulerGamma = 0.5772156649

func harmonicNumber(x float64) float64 {
	return math.Log(x) + eulerGamma
}

func averagePathLength(numSamples int) float64 {
	if numSamples > 2 {
		x := float64(numSamples)
		return 2.0*harmonicNumber(x-1) - 2.0*(x-1)/x
	} else if numSamples == 2 {
		return 1.0
	} else {
		return 0.0
	}
}
