package twoNumberAdd

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{11, 22, 44, 33, 66, 55},
			33,
			[]int{11, 22},
		},
		{
			[]int{1, 2, 4, 3, 6, 5, 8, 7},
			5,
			[]int{1, 4}, //此处23也可以，看题意要求处理
		},
	}
	for _, test := range tests {
		if got := TwoSum2(test.input, test.target); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("ThreeAdd(%v) = %v", test.input, got)
		}
	}
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
		TwoSum3(s, 7)
	}
}

/*func BenchmarkTwoSum4(b *testing.B) {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		TwoSum4(s, 7)
	}
}*/

/*
BenchmarkTwoSum1-8   	20697021	        59.4 ns/op
BenchmarkTwoSum2-8   	33728853	        33.6 ns/op        性能最高
BenchmarkTwoSum3-8   	23456234	        47.0 ns/op
*/
