package _8_2_reverse_left_word

//输入字符串s 和 n，将字符串前k个元素移动到字符串尾部
// 输入abcdef，2，  输出为cdefab

//方案1 利用字符串是可读的byte切片特性，速度快
func reverseLeftWord(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[n:] + s[:n]
}

//方案2 多次string相加导致效率不佳
func reverseLeftWord2(s string, n int) string {
	if len(s) <= n {
		return s
	}
	var str1 string
	var str2 string
	for i := 0; i < len(s); i++ {
		if i < n {
			str1 += string(s[i])
		} else {
			str2 += string(s[i])
		}
	}
	return str2 + str1
}
