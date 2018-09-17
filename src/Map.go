package main

import "fmt"

//多层的map。每一层的map需要再单独make申明一下。
func main() {
	/*m := make(map[int]string) //创建一个key类型为int，value为string的map
	m[1] = "hello"
	m[2] = "kf"
	delete(m, 1) //删除map的值
	fmt.Println(m)*/

	m1:= make(map[string]int)
	m1["hello"]=1
	fmt.Print(m1["hello"])
}
