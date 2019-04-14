package f64

import "math"

const (
	absTol = 1e-8
	relTol = 1.0000000000000001e-05
)

func withinTolerance(a, b float64) bool {
	df := math.Abs(a - b)
	ref := absTol + relTol*math.Abs(b)
	if df > ref {
		return false
	}
	return true
}

func IsEqual(a, b *Mx) bool {
	if a.nr != b.nr || a.nc != b.nc {
		return false
	}

	for i := 0; i < a.len; i++ {
		if !withinTolerance(a.Get(i), b.Get(i)) {
			return false
		}
	}
	return true
}
