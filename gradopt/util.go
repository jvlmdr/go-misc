package gradopt

import (
	"math"
	"math/rand"
)

func normDiff(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("different length")
	}
	var s float64
	for i := range a {
		delta := a[i] - b[i]
		s += delta * delta
	}
	return math.Sqrt(s)
}

func randVec(x []float64) {
	for i := range x {
		x[i] = rand.NormFloat64()
	}
}
