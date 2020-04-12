package main

import "fmt"

func getFun() func() {
	fmt.Println("1")
	return func() {
		fmt.Println("3")
	}
}
func main() {
	defer getFun()()
	fmt.Println("2")
}

// defer getFun()   //返回2，1，内部的3不会执行，defer会丢弃返回值
// 返回1,2,3    解析：真正进入defer的是getFun的返回值，所以1执行后，3就进入的defer函数栈 ，然后执行2，最后执行defer栈的3
// 有点意思！
