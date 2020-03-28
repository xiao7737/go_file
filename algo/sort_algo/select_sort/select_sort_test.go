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
		{[]int{7, 5, 4, 8, 6}, []int{4, 5, 6, 7, 8}},
	}
	for _, test := range tests {
		if got := SelectSort(test.arr); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("InsterSort(%v) = %v", test.arr, got)
		}
	}
}
