package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
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
		fmt.Printf("key:%v value:%v type of value:%T\n", k, v, v) //此处的age是float64
	}

	//方案2
	type Person struct {
		Name string `json:"name" structs:"name"`
		Age  int    `json:"age" structs:"age"`
	}

	u2 := &Person{
		Name: "tom",
		Age:  18,
	}

	m3 := structs.Map(&u2)
	for k, v := range m3 {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}

//json中没有整形和浮点之分，全部转换为float64
//方案1：采用json.Decoder代替json的Unmarshal解决问题
//方案2：利用第三方库structs，原理是反射
