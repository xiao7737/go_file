package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 9, 3}, []int{1, 1, 3, 2}},
	}
	for _, test := range tests {
		if got := BubbleSort(test.arr); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("BubbleSort(%v)=%v", test.arr, got)
		}
	}
}

func ExampleBubbleSort() {
	fmt.Println(BubbleSort([]int{2, 1, 4, 3}))
	fmt.Println(BubbleSort([]int{1, 1, 9, 3}))

	// Output:
	// [1 2 3 4]
	// [1 1 3 9]
}
