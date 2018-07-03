package main

import "fmt"

func main() {
	s1 := make([]int, 6, 10) //slice包含三个元素
	s1[4] = 5
	fmt.Println(len(s1))
	fmt.Println(s1)
	fmt.Println(s1[4])
}
