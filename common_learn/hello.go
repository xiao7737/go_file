package main

import "fmt"

func main() {
	str := "1" + ":" + "" + ":" + "1" // 不能用单引号拼接字符串
	fmt.Println(str)
}
