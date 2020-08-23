package main

import (
	"fmt"
	"runtime"
)

var quit = make(chan int)

func loop() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d.", i)
	}
	quit <- 0
}
func main() {
	// 默认的, 所有goroutine会在一个原生线程里跑，也就是只使用了一个CPU核
	// 最多使用2个核，没有此行的设置，两个go只会在一个线程里面，不是并行计算
	// 在同一个原生线程里，如果当前goroutine不发生阻塞，它是不会让出CPU时间给其他同线程的goroutines的
	// 设置核心数后，才会真正的抢占式运行
	runtime.GOMAXPROCS(runtime.NumCPU())
	go loop()
	go loop()
	for i := 0; i < 2; i++ {
		<-quit
	}
}
