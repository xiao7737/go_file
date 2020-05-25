package pow_50

// 题目：实现 pow(x,n) 计算x的n次方，存在n为负数的情况

// 方案1：暴力法，直接相乘，O(N)
func myPow(x float64, n int) float64 {
	return 0.0
}

// 方案2：分治法，将x的n次方，转换为求2个 x的n/2次方，依次分解成子问题，O(logN)
func MyPow1(x float64, n int) float64 {
	if n >= 0 {
		return helper(x, n)
	}
	// n为负数，取倒数
	return 1.0 / helper(x, -n)
}

// 递归处理，时间复杂度O(logN)，空间复杂度O(logN)
func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// 分解子问题，递归调用
	son := helper(x, n/2)
	// 判断n的奇偶
	if n%2 == 0 {
		return son * son
	}
	return son * son * x
}

// 利用位运算，迭代处理，时间复杂度O(logN)，空间复杂度O(1)
func helper1(x float64, n int) float64 {
	ans := 1.0
	// 贡献的初始值为 x
	x_value := x
	// 在对 N 进行二进制拆分的同时计算答案
	for n > 0 {
		if n%2 == 1 {
			// 如果 N 二进制表示的最低位为 1，，那么需要计入贡献
			ans *= x_value
		}
		// 将贡献不断地平方
		x_value *= x_value
		// 舍弃 n 二进制表示的最低位，这样我们每次只要判断最低位即可
		n /= 2
	}
	return ans
}
