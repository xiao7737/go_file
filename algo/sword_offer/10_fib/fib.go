package main

//返回第n个fib数，该fib从0开始

//时间O(N)，空间(1)
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//dp做法
func fib2(n int) int {
	if n <= 2 {
		return n
	}
	mem := make([]int, n+1)
	mem[0] = 0
	mem[1] = 1
	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}

//fib3  将每个fib保存下来
