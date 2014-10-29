package split

import (
	"fmt"
	"reflect"
)

// Split takes an input of []X and returns an input of [][]X.
func Split(x interface{}, minNum, maxSize int) (interface{}, [][]int) {
	xval := reflect.ValueOf(x)
	n := xval.Len()
	// Split m into the largest groups allowed
	// but do not allow there to be too few groups.
	// Number of groups cannot exceed number of elements.
	m := max(ceilDiv(n, maxSize), min(minNum, n))
	y := reflect.MakeSlice(reflect.SliceOf(xval.Type()), m, m)
	inds := make([][]int, m)
	for i := 0; i < m; i++ {
		yi := reflect.MakeSlice(xval.Type(), 0, ceilDiv(n, m))
		p := make([]int, 0, ceilDiv(n, m))
		for j := 0; m*j+i < n; j++ {
			ind := m*j + i
			yi = reflect.Append(yi, xval.Index(ind))
			p = append(p, ind)
		}
		y.Index(i).Set(yi)
		inds[i] = p
	}
	return y.Interface(), inds
}

//	// SplitTo calls Split.
//	// If x has type []X, then dst must have type *[][]X and be non-nil.
//	func SplitTo(dst, x interface{}, minNum, maxSize int) [][]int {
//		y, inds := Split(x, minNum, maxSize)
//		// De-reference dst pointer.
//		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(y))
//		return inds
//	}

// Takes a slice [][]X and returns a slice []X.
func Merge(x interface{}) interface{} {
	xval := reflect.ValueOf(x)
	m := xval.Len()
	if m == 0 {
		return nil
	}

	p := xval.Index(0).Len()
	y := reflect.MakeSlice(xval.Type().Elem(), 0, m*p)
	for j := 0; j < p; j++ {
		for i := 0; i < m; i++ {
			xi := xval.Index(i)
			if j >= xi.Len() {
				break
			}
			y = reflect.Append(y, xi.Index(j))
		}
	}
	return y.Interface()
}

// Assumes that len(dst) = sum_i len(src[i]).
func MergeTo(dst, src interface{}) {
	srcval := reflect.ValueOf(src)
	dstval := reflect.ValueOf(dst)
	m := srcval.Len()
	// Count number of elements.
	var n int
	for i := 0; i < m; i++ {
		n += srcval.Index(i).Len()
	}
	if dstval.Len() != n {
		panic(fmt.Sprintf("slice len: expect %d, got %d", n, dstval.Len()))
	}
	if m == 0 {
		return
	}
	// Copy elements.
	p := srcval.Index(0).Len()
	ind := 0
	for j := 0; j < p; j++ {
		for i := 0; i < m; i++ {
			xi := srcval.Index(i)
			if j >= xi.Len() {
				break
			}
			dstval.Index(ind).Set(xi.Index(j))
			ind++
		}
	}
}
