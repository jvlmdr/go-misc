package gradopt

import (
	"fmt"
	"math"

	"github.com/gonum/floats"
)

// ConvergenceTest defines an interface for termination criteria.
type ConvergenceTest func(n int, x, p []float64, fx, fp float64, gx []float64) (bool, error)

type SimpleTest struct {
	MaxIter     int
	RelFuncTol  float64
	GradTol     float64
	RelParamTol float64
}

func (crit SimpleTest) Eval(n int, x, p []float64, fx, fp float64, gx []float64) (bool, error) {
	if p == nil {
		return false, nil
	}
	if crit.MaxIter > 0 && n >= crit.MaxIter {
		return false, fmt.Errorf("reached iteration limit: %d", crit.MaxIter)
	}
	if math.Abs(fx-fp) < crit.RelFuncTol*fx {
		return true, nil
	}
	if floats.Norm(gx, math.Inf(1)) < crit.GradTol {
		return true, nil
	}
	if normDiff(x, p) < (floats.Norm(x, 2)+crit.RelParamTol)*crit.RelParamTol {
		return true, nil
	}
	return false, nil
}
