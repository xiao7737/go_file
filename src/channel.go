package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将sum送入管道c
}
func main() {
	a := []int{1, 2, 3, 4}
	c := make(chan int)
	go sum(a[:len(a)/2], c) //送入c的值为3
	go sum(a[len(a)/2:], c) //送入c的值为7
	x, y := <-c, <-c        //从管道中获取
	fmt.Println(x, y)       //7，3
}
