package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64 //装计算好的数组
func main() {
	/*var result uint64= 0  */
	var result = 0
	start := time.Now()
	for i := 1; i < LIM; i++ {
		result = fib(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}

	delta := time.Since(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}

//返回斐波那契数列中指定的第n个数
//性能和fib2()差别不大
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//使用递归:返回斐波那契数列中指定的第n个数  1s
func fib1(n int) (res uint64) {
	if n <= 2 {
		res = 1
	} else {
		res = fib1(n-1) + fib1(n-2)
	}
	return
}

//利用计算机的内存缓存来优化，因为每次计算第n个数都要计算n的前两个数，可以将计算过的值存起来减少重复计算量
//对比fib1()，优势巨大，可以提升一个数量级  0.1s
func fib2(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 2 {
		res = 1
	} else {
		res = fib2(n-1) + fib2(n-2)
	}
	fibs[n] = res
	return
}
