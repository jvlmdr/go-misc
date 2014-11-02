package golden

import "math"

// 0.618
var phi float64 = (math.Sqrt(5) - 1) / 2

type Func func(float64) (float64, error)

func Min(f Func, a, b, tol float64) (float64, error) {
	x := a + phi*(b-a)
	fa, err := f(a)
	if err != nil {
		return 0, err
	}
	fb, err := f(b)
	if err != nil {
		return 0, err
	}
	fx, err := f(x)
	if err != nil {
		return 0, err
	}
	return min(f, a, fa, x, fx, b, fb, tol)
}

func min(f Func, a, fa, x, fx, b, fb, tol float64) (float64, error) {
	m := (a + b) / 2
	if b-a < tol {
		return m, nil
	}
	var (
		p, fp float64
		q, fq float64
		err   error
	)
	if x < m {
		p, fp = x, fx
		q = a + phi*(b-a)
		fq, err = f(q)
		if err != nil {
			return 0, err
		}
	} else {
		q, fq = x, fx
		p = b + phi*(a-b)
		fp, err = f(p)
		if err != nil {
			return 0, err
		}
	}
	if fq < fp {
		return min(f, p, fp, q, fq, b, fb, tol)
	}
	return min(f, a, fa, p, fp, q, fq, tol)
}
