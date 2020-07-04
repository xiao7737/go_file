package _5_replace_space

import "strings"

//替换字符串中的空格
func replaceSpace1(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func replaceSpace2(s string) string {
	var newStr string
	for _, v := range s {
		if v == ' ' {
			newStr += "%20"
		} else {
			newStr += string(v)
		}
	}
	return newStr
}

//利用string.Builder
func replaceSpace3(s string) string {
	var newStr strings.Builder
	for i := range s {
		if s[i] == ' ' {
			newStr.WriteString("%20")
		} else {
			newStr.WriteByte(s[i])
		}
	}
	return newStr.String()
}
