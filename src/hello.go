package main

import "fmt"

func main() {
	fmt.Println("The first is mine！")
	a := 1
	var p *int
	p = &a
	/*var p = &a*/
	fmt.Println(*p)
}
