package _4_find_num_in_2D_arr

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	tests := []struct {
		input  [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 2, 3}, {3, 4, 5}}, 5, true},
		{[][]int{{1, 2, 3}, {3, 4, 5}}, 6, false},
		{[][]int{{5}}, 5, true},
		{[][]int{{}}, 5, false},
	}
	t.Run("FindNumberIn2DArray1", func(t *testing.T) {
		for _, test := range tests {
			if got := FindNumberIn2DArray1(test.input, test.target); got != test.want {
				t.Errorf("got：%v, want:%v", got, test.want)
			}
		}
	})
	t.Run("FindNumberIn2DArray2", func(t *testing.T) {
		for _, test := range tests {
			if got := FindNumberIn2DArray1(test.input, test.target); got != test.want {
				t.Errorf("got：%v, want:%v", got, test.want)
			}
		}
	})
}

func BenchmarkFindRepeatNum(b *testing.B) {
	b.Run("FindRepeatNum1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindNumberIn2DArray1([][]int{{1, 2, 3}, {3, 4, 5}}, 5)
		}
	})
	b.Run("FindRepeatNum2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindNumberIn2DArray2([][]int{{1, 2, 3}, {3, 4, 5}}, 5)
		}
	})
}
