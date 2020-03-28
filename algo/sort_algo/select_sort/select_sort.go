package main

// 选择排序 不稳定 时间复杂度 O(n*n)，空间O(1)
// 思路：第一次从未排序数组中找到最小的，放到第一个位置，第二次从未排序数组中找到最小的，放到第二个位置
func SelectSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}
	for i := 0; i < arrLen; i++ {
		// 用tail存已排序的最后一个元素
		tail := i
		for j := i + 1; j < arrLen; j++ {
			// 找到未排序元素中最小的元素
			if arr[j] < arr[tail] {
				tail = j
			}
		}
		// 如果tail发生了改变，则交换
		if tail != i {
			arr[tail], arr[i] = arr[i], arr[tail]
		}

	}
	return arr
}
