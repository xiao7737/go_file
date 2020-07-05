package _6_pow

//求一个数的整数次方，考虑负数的情况
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return helper(x, n)
	} else {
		return 1.0 / helper(x, -n)
	}
}

// 递归处理，时间复杂度O(logN)，空间复杂度O(logN)
func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	sub := helper(x, n/2)
	if n%2 == 0 {
		return sub * sub
	} else {
		return sub * sub * x
	}
}
