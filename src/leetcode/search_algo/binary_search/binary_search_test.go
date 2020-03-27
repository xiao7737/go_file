package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		sortedArr []int
		searchKey int
		want      int
	}{
		{[]int{1, 3, 4, 6, 10, 13}, 6, 3},
		{[]int{1, 3, 4, 6, 10, 13}, 5, -1},
	}
	for _, test := range tests {
		if got := BinarySearch(test.sortedArr, test.searchKey); got != test.want {
			t.Errorf("BinarySearch(%v, %d) = %d", test.sortedArr, test.searchKey, got)
		}
	}
}

// 查找第一个searchKey出现的元素
func TestBinarySearch1(t *testing.T) {
	var tests = []struct {
		sortedArr []int
		searchKey int
		want      int
	}{
		{[]int{1, 3, 4, 6, 6, 13}, 6, 3},
		{[]int{1, 3, 4, 4, 4, 13}, 4, 2},
	}
	for _, test := range tests {
		if got := BinarySearch1(test.sortedArr, test.searchKey); got != test.want {
			t.Errorf("BinarySearch1(%v, %d) = %d", test.sortedArr, test.searchKey, got)
		}
	}
}

// 查找第一个比searchKey大的元素
func TestBinarySearch2(t *testing.T) {
	tests := []struct {
		sortedArr []int
		searchKey int
		want      int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 3},
	}
	for _, test := range tests {
		if got := BinarySearch2(test.sortedArr, test.searchKey); got != test.want {
			t.Errorf("BinarySearch2(%v, %d) = %d", test.sortedArr, test.searchKey, got)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch([]int{1, 3, 4, 6, 10, 13}, 6)
	}
}
