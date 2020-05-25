package pow_50

import "testing"

func TestMyPow1(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{2, 10, 1024},
		{2.0, -2, 0.25},
	}
	for _, test := range tests {
		if got := MyPow1(test.x, test.n); got != test.want {
			t.Errorf("MyPow(%v,%v)=%v", test.x, test.n, got)
		}
	}
}
