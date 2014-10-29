package split_test

import (
	"reflect"
	"testing"

	"github.com/jvlmdr/go-misc/split"
)

var splitTests = []struct {
	MinNum  int
	MaxSize int
	In      interface{}
	Out     interface{}
}{
	{
		1, 3,
		[]int{1, 2, 3, 4, 5, 6},
		[][]int{{1, 3, 5}, {2, 4, 6}},
	},
	{
		1, 3,
		[]string{"one", "two", "three", "four", "five", "six"},
		[][]string{{"one", "three", "five"}, {"two", "four", "six"}},
	},
	{
		1, 2,
		[]int{1, 2, 3, 4, 5, 6},
		[][]int{{1, 4}, {2, 5}, {3, 6}},
	},
	{
		1, 3,
		[]int{1, 2, 3, 4, 5},
		[][]int{{1, 3, 5}, {2, 4}},
	},
	{
		1, 3,
		[]int{1, 2, 3, 4, 5, 6, 7},
		[][]int{{1, 4, 7}, {2, 5}, {3, 6}},
	},
	{
		3, 3,
		[]int{1, 2, 3, 4, 5, 6, 7},
		[][]int{{1, 4, 7}, {2, 5}, {3, 6}},
	},
	{
		4, 3,
		[]int{1, 2, 3, 4, 5, 6, 7},
		[][]int{{1, 5}, {2, 6}, {3, 7}, {4}},
	},
}

func TestSplit(t *testing.T) {
	for _, x := range splitTests {
		got, _ := split.Split(x.In, x.MinNum, x.MaxSize)
		if !reflect.DeepEqual(x.Out, got) {
			t.Errorf("%+v: got %v", x, got)
		}
	}
}

func TestMerge_AfterSplit(t *testing.T) {
	for _, x := range splitTests {
		y, _ := split.Split(x.In, x.MinNum, x.MaxSize)
		got := split.Merge(y)
		if !reflect.DeepEqual(x.In, got) {
			t.Errorf("%+v: got %v", x, got)
		}
	}
}

func TestMergeTo_AfterSplit(t *testing.T) {
	for _, x := range splitTests {
		y, _ := split.Split(x.In, x.MinNum, x.MaxSize)
		val := reflect.ValueOf(x.In)
		got := reflect.MakeSlice(val.Type(), val.Len(), val.Len()).Interface()
		split.MergeTo(got, y)
		if !reflect.DeepEqual(x.In, got) {
			t.Errorf("%+v: got %v", x, got)
		}
	}
}
