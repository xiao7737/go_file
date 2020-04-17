package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u1 := map[string]interface{}{
		"Name": "tom",
		"Age":  18,
	}
	for k, v := range u1 {
		fmt.Printf("key:%v value:%v type of value:%T\n", k, v, v) //此处age是int
	}
	b, _ := json.Marshal(u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v type of value:%T\n", k, v, v)
		//此处的age变成了float64，小坑
		//json中没有整形和浮点之分，全部转换为float64
		//采用json.Decoder代替jsonUnmarshal解决问题
	}
}
