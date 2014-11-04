package gradopt

import (
	"fmt"
	"math"
	"os"

	"github.com/gonum/floats"
)

type Func func(x []float64) (fx float64, gx []float64, err error)

func clone(x []float64) []float64 {
	y := make([]float64, len(x))
	copy(y, x)
	return y
}

const minLipschitz = 1e-6

// Desc performs gradient descent on a smooth function.
func Desc(f Func, x []float64, test ConvergenceTest) ([]float64, error) {
	l, err := lowerBoundLipschitz(f, x)
	if err != nil {
		return nil, fmt.Errorf("bound Lipschitz constant: %v", err)
	}
	if l < minLipschitz {
		l = minLipschitz
	}
	return DescRate(f, x, test, 1/l)
}

func DescRate(f Func, x []float64, test ConvergenceTest, t float64) ([]float64, error) {
	var (
		p  []float64
		fp float64
	)
	for k := 0; ; k++ {
		fx, gx, err := f(x)
		if err != nil {
			return nil, err
		}
		if math.IsNaN(fx) {
			return nil, fmt.Errorf("function evaluates to NaN")
		}
		if floats.HasNaN(gx) {
			return nil, fmt.Errorf("gradient evaluates to NaN")
		}
		if k > 0 {
			fmt.Fprintf(
				os.Stderr, "%5d  f:%13.6e  df:%10.3e  g:%10.3e  dx:%10.3e  t:%10.3e\n",
				k, fx, fx-fp, floats.Norm(gx, math.Inf(1)), normDiff(x, p), t,
			)
		}
		conv, err := test(k, x, p, fx, fp, gx)
		if err != nil {
			return nil, err
		}
		if conv {
			return x, nil
		}
		p, fp = x, fx
		z := make([]float64, len(x))
		for satisfied := false; !satisfied; {
			// Tentatively consider x <- x - t gx.
			floats.AddScaledTo(z, x, -t, gx)
			// TODO: Avoid unnecessary gradient computation.
			fz, _, err := f(z)
			if err != nil {
				return nil, err
			}
			if math.IsNaN(fz) {
				return nil, fmt.Errorf("backtrack: function evaluates to NaN")
			}
			// Need to ensure that t satisfies
			//   fx-fz >= 0.5*t*floats.Dot(gx, gx).
			// If we decrease the step size too much, the updates vanish.
			// With tolerance, the left hand side is
			//   (fx+a) - (fz+b)
			// where |a| <= eps |fx|, |b| <= eps |fz|
			//   = (fx - fz) + (a - b)
			// where |a - b| <= eps (|fx| + |fz|).
			// TODO: Account for gradient?
			// The gradient term is dot(gx+c, gx+c) with |c_i| <= eps |gx_i|.
			//   dot(gx+c, gx+c) = dot(gx+c) + 2*dot(gx, c)
			// and
			//   2*dot(gx, c) <= 2*eps*norm(gx, 1)
			const eps = 1e-12
			delta := eps * (math.Abs(fx) + math.Abs(fz))
			if fx-fz+delta >= 0.5*t*floats.Dot(gx, gx) {
				satisfied = true
			} else {
				t /= 2
			}
		}
		x = z
	}
}

// Uses a random perturbation to bound the Lipschitz constant.
func lowerBoundLipschitz(f Func, x []float64) (float64, error) {
	// ||f'(x) - f'(y)|| <= L ||x - y||
	// L >= ||f'(x) - f'(y)|| / ||x - y||
	_, gx, err := f(x)
	if err != nil {
		return 0, err
	}
	if floats.HasNaN(gx) {
		return 0, fmt.Errorf("gradient evaluates to NaN")
	}
	y := make([]float64, len(x))
	randVec(y)
	floats.Add(y, x)
	_, gy, err := f(y)
	if err != nil {
		return 0, err
	}
	return normDiff(gx, gy) / normDiff(x, y), nil
}
