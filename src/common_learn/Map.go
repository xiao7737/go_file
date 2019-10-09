package main

import (
	"fmt"
	"sort"
)

//多层的map。每一层的map需要再单独make申明一下。
func main() {
	/*m := make(map[int]string) //创建一个key类型为int，value为string的map
	m[1] = "hello"
	m[2] = "go"
	delete(m, 1)   //删除map的值
	fmt.Println(m) //map[2:kf]

	m1 := make(map[string]int)
	m1["hello"] = 1
	m1["go"] = 2

	for k, v := range m1 {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println("The value:", m[2])
	fmt.Println("The value:", m1["hello"])

	//双键检测，如果存在，ok的值为true，否则为true
	elem, ok := m1["hello"]
	if ok {
		fmt.Println("The value:", elem)
	}*/

	//实现map的顺序读取:将map的key读入slice，对slice中的key排序
	map1 := map[string]int{
		"apple":  1,
		"beer":   2,
		"orange": 3,
	}

	//put the key in a slice and sort it
	var keys []string
	for key := range map1 {
		keys = append(keys, key)
	}
	sort.Strings(keys) //此处采用字母排序

	//display the key according the sorted slice
	for _, key := range keys {
		fmt.Printf("%s:%v\n", key, map1[key])
	}
}
