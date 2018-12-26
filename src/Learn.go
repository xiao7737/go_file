package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
}
