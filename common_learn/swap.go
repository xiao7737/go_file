package main

import "fmt"

//多值返回
func swap(x, y string) (a, b string) {
	b, a = x, y
	return
}
func main() {
	fmt.Println(swap("hello", "go"))
}
