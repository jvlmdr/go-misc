package gradopt_test

import (
	"testing"

	"github.com/gonum/floats"
	"github.com/jvlmdr/go-misc/gradopt"
)

func TestMin(t *testing.T) {
	cases := []struct {
		F    gradopt.Func
		X    []float64
		Test gradopt.ConvergenceTest
		Want []float64
	}{
		{
			F: func(x []float64) (float64, []float64, error) {
				f := 3*x[0]*x[0] + 2*x[0]*x[1] + x[1]*x[1]
				g := []float64{6*x[0] + 2*x[1], 2*x[0] + 2*x[1]}
				return f, g, nil
			},
			X:    []float64{1, 1},
			Test: gradopt.SimpleTest{GradTol: 1e-6}.Eval,
			Want: []float64{0, 0},
		},
		{
			F: func(x []float64) (float64, []float64, error) {
				f := x[0]*x[0] - 4*x[0] + 2*x[0]*x[1] + 2*x[1]*x[1] + 2*x[1] + 14
				g := []float64{2*x[0] - 4 + 2*x[1], 2*x[0] + 4*x[1] + 2}
				return f, g, nil
			},
			X:    []float64{4, -4},
			Test: gradopt.SimpleTest{GradTol: 1e-9}.Eval,
			Want: []float64{5, -3},
		},
	}

	for _, p := range cases {
		got, err := gradopt.Desc(p.F, p.X, p.Test)
		if err != nil {
			t.Error(err)
			continue
		}
		if !floats.EqualApprox(p.Want, got, 1e-6) {
			t.Errorf("want %v, got %v", p.Want, got)
		}
	}
}
