package main

import (
	"fmt"
	"time"
)

/*func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将sum送入管道c
}
func main() {
	a := []int{1, 2, 3, 4}
	c := make(chan int)
	go sum(a[:len(a)/2], c) //送入c的值为3=1+2
	go sum(a[len(a)/2:], c) //送入c的值为7=3+4
	x, y := <-c, <-c        //从管道中获取
	fmt.Println(x, y)       //7，3
}*/

//用channel select实现斐波纳契数据
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
		case <-time.After(time.Millisecond):
			fmt.Println("time out")

		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

//channel源码解析
//channel底层数据是hchan struct
