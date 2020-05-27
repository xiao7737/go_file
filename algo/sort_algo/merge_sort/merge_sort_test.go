package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{2, 3, 1, 4}, []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		if got := MergeSort(test.arr); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("MergeSort(%v) = %v", test.arr, got)
		}
	}
}
