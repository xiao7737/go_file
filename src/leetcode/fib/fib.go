package fib

//返回斐波那契数列中指定的第n个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

//使用递归:返回斐波那契数列中指定的第n个数
func fib1(n int) (res int) {
	if n < 2 {
		res = 1
	} else {
		res = fib1(n-1) + fib1(n-2)
	}
	return
}
