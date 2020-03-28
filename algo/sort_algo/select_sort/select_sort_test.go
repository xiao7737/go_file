package main

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{2, 3, 1, 4}, []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		if got := SelectSort(test.arr); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("InsterSort(%v) = %v", test.arr, got)
		}
	}
}
