package main

// 选择排序 不稳定
// 时间复杂度 O(n*n)，空间O(1)
func SelectSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		min := i //arr[i]是已排序数组中最小的
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
