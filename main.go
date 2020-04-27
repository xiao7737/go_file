package main

func main() {
	x := 1
	println(x) // 1
	{
		println(x) // 1
		x := 2
		println(x) // 2    // 新的 x 变量的作用域只在代码块内部
	}
	println(x) // 1
}
