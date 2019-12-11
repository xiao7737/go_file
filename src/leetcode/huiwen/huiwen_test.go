package huiwen

import (
	"testing"
)

// 功能测试：go test go_file/src/common_learn/huiwen    从GOPATH路径开始写
// 覆盖率：go test -cover
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"Aba", true},
		{"aab，bAA", true},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q)=%v", test.input, got)
		}
	}
}

// 性能测试：go test -bench=IsPalindrome
// 也可以 go test -bench=.   匹配文件中所有的bench测试函数
// go test -bench=. -benchmem    查看运行过程中，内存分配的次数，目前适合三次
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("aab,bAA")
	}
}

func TestIsIsPalindrome1(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{-1, false},
		{1, true},
		{121, true},
		{133, false},
	}
	for _, test := range tests {
		if got := IsIsPalindrome1(test.input); got != test.want {
			t.Errorf("IsIsPalindrome1(%q)=%v", test.input, got)
		}
	}
}
