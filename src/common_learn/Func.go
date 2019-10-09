package main

//匿名函数和闭包，即函数像普通变量一样被传递或使用，匿名函数不能作为最外层函数
/*func main() {
	v := func(a int) int {
		return a * a
	}
	fmt.Println(v(3))
}*/

/*func main() {
	a := 1
	A(&a)
	fmt.Println(a)
}
func A(a *int) {
	*a = 2
	fmt.Println(*a)
} //int.string为值类型，为值拷贝，操作的只是副本，用指针，改变原有值。结果为2,2*/

/*
func main() {
	a := A
	a()
}
func A() {
	fmt.Println("func A")         //函数也是一种类型，可以传递
}*/

/*func main() {
	A()
	B()
}
func A() {
	fmt.Println("func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("Panic B")
	//panic可在任何位置引发，recover只能在defer中使用匿名函数调用,
	//recover在panic之前用,恢复运行
}*/

/*func main() {     //利用指针交换变量值
	a := 1
	b := 2
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
func swap(x, y *int) {
	*x, *y = *y, *x
}*/

/*func main() { //直接交换变量
	a := 1
	b := 22
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}
func swap(x, y int) (int, int) {
	y, x = x, y
	return x, y
}*/

import "fmt"

// 返回a、b中最大值.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) //调用函数max(x, y)
	max_xz := max(x, z) //调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它
}
