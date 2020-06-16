package sqrt_69

import (
	"math"
)

// 给定x，求出x的平凡根，给定4，返回2
// 只返回整数部分

// 二分法
func mySqrt(x int) (res int) {
	if x == 0 || x == 1 {
		return x
	}
	low := 0
	high := x
	for low <= high {
		mid := (low + high) >> 1
		if mid*mid > high {
			high = mid
		} else if mid*mid < high {
			low = mid
		} else if math.Abs(float64(high-low)) > 1e-5 {
			return mid
		}
	}
	return -1
}

// 方法一样，改良版
func mySqrt1(x int) (res int) {
	if x == 0 || x == 1 {
		return x
	}
	low := 0
	high := x
	for low <= high {
		mid := low + (high-low)>>1 //一样的效果，避免low+high过大
		if mid > high/mid {        //一样的效果，避免mid*mid过大
			high = mid
		} else if mid < high/mid {
			low = mid
		} else if math.Abs(float64(high-low)) > 1e-5 { // 精度
			return mid
		}
	}
	return -1
}

//只返回整数部分
func mySqrt2(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	var res int
	low, high := 0, x
	for low <= high {
		mid := (low + high) >> 1
		if mid > x/mid {
			high = mid - 1
		} else if mid <= x/mid {
			low = mid + 1
			res = mid
		}
	}
	return res
}
