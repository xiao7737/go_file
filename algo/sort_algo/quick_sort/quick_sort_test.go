package main

import (
	"reflect"
	"testing"
)

func TestInsertSort1(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{11, 22, 44, 33, 66, 55}, []int{11, 22, 33, 44, 55, 66}},
	}
	for _, test := range tests {
		if got := QuickSort1(test.arr); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("InsterSort(%v) = %v", test.arr, got)
		}
	}
}
func TestInsertSort2(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{11, 22, 44, 33, 66, 55}, []int{11, 22, 33, 44, 55, 66}},
	}
	for _, test := range tests {
		if got := QuickSort2(test.arr); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("InsterSort(%v) = %v", test.arr, got)
		}
	}
}

var testSli = []int{1, 43, 54, 62, 21, 8, 32}

func BenchmarkQuickSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort1(testSli)
	}
}
func BenchmarkQuickSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort2(testSli)
	}
}
func BenchmarkQuickSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort3(testSli)
	}
}
