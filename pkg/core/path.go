package core

import "math"

const eulerGamma = 0.5772156649

func harmonicNumber(x float64) float64 {
	return math.Log(x) + eulerGamma
}

func averagePathLength(x int) float64 {
	if x > 2 {
		f := float64(x)
		return 2.0*harmonicNumber(f-1) - 2.0*(f-1)/f
	} else if x == 2 {
		return 1.0
	} else {
		return 0.0
	}
}
