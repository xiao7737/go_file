package twoNumberAdd

import (
	"testing"
)

func TestTwoSum1(t *testing.T) {

}
func TestTwoSum2(t *testing.T) {

}

func BenchmarkTwoSum1(b *testing.B) {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		TwoSum1(s, 7)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		TwoSum2(s, 10)
	}
}

func BenchmarkTwoSum3(b *testing.B) {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		TwoNum3(s, 7)
	}
}

/*func BenchmarkTwoSum4(b *testing.B) {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		TwoNum4(s, 7)
	}
}*/

/*
BenchmarkTwoSum1-8   	20697021	        59.4 ns/op
BenchmarkTwoSum2-8   	33728853	        33.6 ns/op        性能最高
BenchmarkTwoSum3-8   	23456234	        47.0 ns/op
*/
