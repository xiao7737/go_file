package main

import "fmt"

func main() {
	a := 1
	var p *int
	p = &a
	str := "123" + "abc" //不能用单引号拼接字符串
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(str)
}
