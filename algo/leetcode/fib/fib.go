package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]int //装计算好的数组
func main() {
	var result = 0
	start := time.Now()
	for i := 1; i < LIM; i++ {
		result = fib3(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	run := time.Since(start)
	fmt.Printf("程序的执行时间为: %s\n", run)
}

//返回斐波那契数列中指定的第n个数
//时间O(N)，空间O(1)
//性能和fib2()差别不大
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//使用递归:返回斐波那契数列中指定的第n个数  1s
//时间复杂度：2的N次方
func fib1(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fib1(n-1) + fib1(n-2)
	}
	return
}

//利用计算机的内存缓存来优化，因为每次计算第n个数都要计算n的前两个数，可以将计算过的值存起来减少重复计算量
//对比fib1()，优势巨大，可以提升一个数量级  0.1s
func fib2(n int) (res int) {
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

//动态规划
//时间O(N)，空间O(n)
//优化：可以采用fib第一种解法，将数组用两个变量替换，空间变成O(1)
func fib3(n int) int {
	if n <= 2 {
		return 1
	}
	mem := make([]int, n)
	mem[0] = 1
	mem[1] = 1
	for i := 2; i < n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n-1]
}

func fib4(n int) int {
	var s []int
	ch := make(chan int, 40)
	helper(n, ch)
	for i := range ch {
		s = append(s, i)
	}
	return s[len(s)-1]
}

func helper(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		ch <- x
	}
	close(ch)
}
