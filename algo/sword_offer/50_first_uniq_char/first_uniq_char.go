package _0_first_uniq_char

//在字符串 s 中找出第一个只出现一次的字符，如果没有，返回一个单空格。 s只包含小写字母

//利用hash表，key是元素，value是元素出现的次数
//需要扫描两次，扫描第一次，进行计数，第二次，返回次数为1的元素，总的时间O(N)
//注意go的map遍历返回时无序的，用slice保证顺序
func firstUniqChar(s string) byte {
	l := len(s)
	if l == 0 {
		return byte(' ')
	}
	bytes := []byte(s)
	m := make(map[byte]int, len(s))
	sli := make([]byte, 0)

	for _, v := range bytes {
		if _, ok := m[v]; ok {
			//map中已经存在该元素
			m[v] += 1
		} else {
			//map中不存在该元素
			m[v] = 1
			sli = append(sli, v)
		}
	}

	for _, v := range sli {
		if m[v] == 1 {
			return v
		}
	}
	return byte(' ')
}
