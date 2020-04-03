package huiwen

import (
	"fmt"
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
	b.Logf("something to show")
	b.ResetTimer() //重置测试时间，可以去除数据准备期间的时间消耗
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

// 在godoc里面生成示例，也能达到测试函数的效果，但是可以有实例，显示比较清晰
func ExampleIsPalindrome() {
	fmt.Println(IsIsPalindrome1(-1))
	fmt.Println(IsIsPalindrome1(1))
	fmt.Println(IsIsPalindrome1(121))
	fmt.Println(IsIsPalindrome1(133))

	// Output:
	// false
	// true
	// true
	// false
}

//	 go test -bench . -benchmem -cpuprofile prof.cpu   获取CPU性能数据
//        go tool pprof prof.cpu
//go test -bench . -benchmem -memprofile prof.mem    获取内存性能数据
// 		  go tool pprof prof.men

//  pprof 命令下
// 输入:list IsPalindrome 查看具体耗时
// 输入：web  查看展示图
