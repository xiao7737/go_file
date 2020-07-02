package dp_length_of_lis_300

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 3, 1, 4}, 3},
		{[]int{2, 3, 4, 4}, 3},
		{[]int{2, 1, 1, 0}, 1},
	}
	for _, test := range tests {
		if got := LengthOfLIS(test.arr); test.want != got {
			t.Errorf("InsterSort(%v) = %v", test.arr, got)
		}
	}
}
