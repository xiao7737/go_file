package _0_min_kth_num

import (
	"sort"
)

//题目：返回最小的k个数

//方案1：排序后直接返回
//时间O(N*logN)
//空间O(logN)
func getLeastNumbers(arr []int, k int) []int {
	var res []int
	sort.Ints(arr)
	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}
	return res
}

//利用堆
func getLeastNumbers1(arr []int, k int) []int {

}

//利用快排思想
func getLeastNumbers2(arr []int, k int) []int {

}
