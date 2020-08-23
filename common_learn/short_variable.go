package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(&a)
	a, b := 2, 3
	fmt.Println(a, b)
	fmt.Println(&a)
}

//两次a的是一个地址，只是被重新赋值，并没有重新申明
// https://golang.org/ref/spec#
