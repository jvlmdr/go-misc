package golden_test

import (
	"math"
	"testing"

	"github.com/jvlmdr/go-misc/golden"
)

func TestMin(t *testing.T) {
	cases := []struct {
		F    golden.Func
		A, B float64
		Tol  float64
		Min  float64
	}{
		{
			F:   func(x float64) (float64, error) { return math.Log(math.Pow(x-1, 2) + 1e-3), nil },
			A:   0,
			B:   100,
			Tol: 1e-6,
			Min: 1,
		},
		{
			F:   func(x float64) (float64, error) { return x, nil },
			A:   -1,
			B:   1,
			Tol: 1e-6,
			Min: -1,
		},
	}

	for _, p := range cases {
		got, err := golden.Min(p.F, p.A, p.B, p.Tol)
		if err != nil {
			t.Errorf("error in func: %v", err)
			continue
		}
		if math.Abs(got-p.Min) > p.Tol {
			t.Errorf("want %g, got %g", p.Min, got)
		}
	}
}
