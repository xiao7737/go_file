package slice

import (
	"reflect"
	"testing"
)

// TestUnion 并集
func TestUnion(t *testing.T) {
	tests := []struct {
		slice1 []string
		slice2 []string
		want   []string
	}{
		{[]string{"1", "2", "3", "4"}, []string{"2", "3", "4", "5"}, []string{"1", "2", "3", "4", "5"}},
	}
	for _, test := range tests {
		if got := Union(test.slice1, test.slice2); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("Union(%v,%v) = %v", test.slice1, test.slice2, got)
		}
	}
}

// TestIntersect 交集
func TestIntersect(t *testing.T) {
	tests := []struct {
		slice1 []string
		slice2 []string
		want   []string
	}{
		{[]string{"1", "2", "3", "4"}, []string{"2", "3", "4", "5"}, []string{"2", "3", "4"}},
	}
	for _, test := range tests {
		if got := Intersect(test.slice1, test.slice2); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("Intersect(%v,%v) = %v", test.slice1, test.slice2, got)
		}
	}
}

// TestDifference 差集
func TestDifference(t *testing.T) {
	tests := []struct {
		slice1 []string
		slice2 []string
		want   []string
	}{
		{[]string{"1", "2", "3", "4"}, []string{"2", "3", "4", "5"}, []string{"1"}}, // 此处只会返回1
	}
	for _, test := range tests {
		if got := Difference(test.slice1, test.slice2); reflect.DeepEqual(test.want, got) == false {
			t.Errorf("Difference(%v,%v) = %v", test.slice1, test.slice2, got)
		}
	}
}
