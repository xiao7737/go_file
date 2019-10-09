package main

import "fmt"

func calc(index string, a, b int) int {
	sum := a + b
	fmt.Println(index, a, b, sum)
	return sum
}
func main() {
	a := 1
	b := 2
	defer /*④*/ calc("1", a /*①*/, calc("10", a, b))
	a = 0
	defer /*③*/ calc("2", a /*②*/, calc("20", a, b))
	b = 1
}

//1234执行顺序
/*结果
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4*/

//延迟函数会被压入栈，按照先进后出的顺序调用
/*
for i := 0; i < 5; i++ {
defer fmt.Println(i)
}
结果为4，3，2，1，0
*/
