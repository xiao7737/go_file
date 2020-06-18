package main

import "fmt"

type T1 struct {
}

func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1 //此处取别名后，T2也拥有T1的所有方法
type MyStruct struct {
	T1
	T2
}

func main() {
	my := MyStruct{}
	my.m1()
	//不能运行，提示 ambiguous selector my.m1，因为此处程序不知道运行T1还是T2的m1方法
	//  改为my.T1.m1()
}
