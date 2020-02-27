package main

import "fmt"

//打印99乘法表
func main() {
	//先遍历行
	for y := 1; y <= 9; y++ {
		//再遍历列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d,", x, y, x*y)
		}
		//换行
		fmt.Println()
	}
}
