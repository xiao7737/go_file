package _7_print_number

import "math"

//输出从1到n位数
//如果n为2，则输出1-99

//找到最大值
func printNumbers(n int) []int {
	if n <= 0 {
		return nil
	}
	var max int
	var res []int
	for n > 0 {
		max = 10*max + 9
		n--
	}
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res
}

//最大值为10的N次方-1
func printNumbers1(n int) []int {
	if n <= 0 {
		return nil
	}

	var res []int
	for i := 1; i < int(math.Pow10(n)); i++ {
		res = append(res, i)
	}
	return res
}
