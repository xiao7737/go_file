package _3

import (
	"testing"
)

func TestFindRepeatNum(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2, 3, 4}, 1},
		{[]int{5, 6, 7, 4, 3, 2, 2}, 2},
		{[]int{1}, -1},
		{[]int{}, -1},
		{[]int{1, 2, 3, 4, 5, 6}, -1},
	}
	t.Run("FindRepeatNum1", func(t *testing.T) {
		for _, test := range tests {
			if got := FindRepeatNum1(test.input); got != test.want {
				t.Errorf("got：%v, want:%v", got, test.want)
			}
		}
	})
	t.Run("FindRepeatNum2", func(t *testing.T) {
		for _, test := range tests {
			if got := FindRepeatNum2(test.input); got != test.want {
				t.Errorf("got：%v, want:%v", got, test.want)
			}
		}
	})
}

func BenchmarkFindRepeatNum(b *testing.B) {
	b.Run("FindRepeatNum1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindRepeatNum1([]int{1, 3, 4, 6, 10, 13})
		}
	})
	b.Run("FindRepeatNum2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindRepeatNum2([]int{1, 3, 4, 6, 10, 13})
		}
	})
}
