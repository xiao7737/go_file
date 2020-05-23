package threeNumberAdd_15

import (
	"reflect"
	"testing"
)

func TestThreeAdd(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{11, 22, 44, 33, 66, 55},
			66,
			[]int{11, 22, 33},
		},
		{
			[]int{1, 2, 4, 3, 6, 5, 8, 7},
			10,
			[]int{1, 2, 7},
		},
	}
	for _, test := range tests {
		if got := ThreeAdd2(test.input, test.target); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("ThreeAdd(%v) = %v", test.input, got)
		}
	}
}

func TestThreeAdd3(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   [][]int
	}{
		{
			[]int{11, 22, 44, 33, 66, 55},
			66,
			[][]int{{11, 22, 33}},
		},
		{
			[]int{0, 0, 0, 0, 1, 2, 4, 3, 6, 5, 6, 2, 8, 7},
			10,
			[][]int{{1, 2, 7}, {1, 3, 6}},
		},
	}
	for _, test := range tests {
		if got := ThreeAdd3(test.input, test.target); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("ThreeAdd(%v) = %v", test.input, got)
		}
	}
}
