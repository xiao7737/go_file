package main

import "fmt"

//func calc(index string, a, b int) int {
//	sum := a + b
//	fmt.Println(index, a, b, sum)
//	return sum
//}
//func main() {
//	a := 1
//	b := 2
//	defer /*④*/ calc("1", a /*①*/, calc("10", a, b))
//	a = 0
//	defer /*③*/ calc("2", a /*②*/, calc("20", a, b))
//	b = 1
//}

//1234执行顺序
/*结果
10 1 2 3  第一个defer执行入栈的时候，发现参数是一个函数，执行该函数后，将defer入栈
20 0 2 2
2 0 2 2
1 1 3 4 此处为最先入栈的defer，最后执行
*/

func main() {
	call()
}
func call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

/*
打印后
打印前
panic: 触发异常  panic会在defer执行完才抛出*/

//延迟函数会被压入栈，按照先进后出的顺序调用
/*
for i := 0; i < 5; i++ {
defer fmt.Println(i)
}
结果为4，3，2，1，0
*/
