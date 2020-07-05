package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4}
	n2 := n[:len(n)-1]
	println(n[len(n)-1])
	fmt.Println(n2)

}
