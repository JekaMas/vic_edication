package sqrt

import "math"

const Precision = 1E-12

func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}

	var z float64
	switch {
	case x < 1:
		z = 1 - Precision
	default:
		z = x
	}

	for n := 0; math.Abs(z*z - x) > Precision; n++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}
