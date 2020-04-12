package main

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

/*func main() {
	call()
}
func call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}*/

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

func main() {
	println(deferFun1(1)) //4
	println(deferFun2(1)) //1
}

func deferFun1(i int) (t int) { //	defer在函数结束前执行，函数返回值t的作用域为整个函数，t会被defer修改
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func deferFun2(i int) int { //  t的作用域仅为闭包，返回的t为最开始赋值的t
	t := i
	defer func() {
		t += 3
	}()
	return t
}
