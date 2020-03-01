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

//利用计算机的内存缓存来优化，因为每次计算第n个数都要计算n的前两个数，可以将计算过的值存起来减少重复计算量
var fibs [10]int //装计算好的数组

func fib2(n int) (res int) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n < 2 {
		res = 1
	} else {
		res = fib1(n-1) + fib1(n-2)
	}
	fibs[n] = res
	return
}
