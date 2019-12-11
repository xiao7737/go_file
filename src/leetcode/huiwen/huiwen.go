package huiwen

import (
	"strconv"
	"unicode"
)

//判断字符串是否为回文，忽略大小写
func IsPalindrome(str string) bool {
	// var strings []int32
	// 改进为以下写法，创建一个合适长度的slice，减少内存分配的次数，能显著提高性能
	strings := make([]int32, 0, len(str))
	for _, v := range str {
		if unicode.IsLetter(v) {
			strings = append(strings, unicode.ToLower(v)) //转换为小写
		}
	}
	for i := range strings {
		if strings[i] != strings[len(strings)-i-1] {
			return false
		}
	}
	return true
}

//判断整数是否是回文
func IsIsPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	y := strconv.Itoa(x)
	for i := range y {
		if y[i] != y[len(y)-i-1] {
			return false
		}
	}
	return true
}
