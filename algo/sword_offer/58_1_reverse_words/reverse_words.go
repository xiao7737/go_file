package _8_1_reverse_words

import "strings"

//翻转句子中单次的顺序，单次内字符顺序不变
//输入: "the sky is blue"
//输出: "blue is sky the"

func reverseWords(s string) string {
	//用空格将s切开
	strArr := strings.Split(s, " ")
	var res []string
	//从最后开始遍历
	for i := len(strArr) - 1; i >= 0; i-- {
		//处理空格
		if len(strArr[i]) > 0 {
			res = append(res, strArr[i])
		}
	}
	//用空格连接
	return strings.Join(res, " ")
}
