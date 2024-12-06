package core

import "math"

const euler = 0.5772156649

func harmonicNumber(x float64) float64 {
	return math.Log(x) + euler
}

func cInt(x int) float64 {
	return c(float64(x))
}
func c(x float64) float64 {
	if x > 2 {
		return 2.0*harmonicNumber(x-1) - 2.0*(x-1)/x
	} else if x == 2 {
		return 1.0
	} else {
		return 0.0
	}
}
