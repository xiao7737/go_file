package mergeTwoArray_88

import "sort"

// 合并两个有序数组

// 方法1  直接将arr2添加到arr1中，然后再对arr1排序
// 没有利用arr的有序性
func mergeTwoArr(arr1, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	return arr1
}

// 方法2  利用双指针，类似归并算法的merge
func mergeTwoArr1(arr1, arr2 []int) []int {
	i, j := 0, 0
	var c []int
	// 将arr1和arr2中的元素进行比较，将小的元素依次装入合并空间
	// 直到某一个先遍历完，则将另一个剩下的元素直接复制到合并序列的尾部
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			c = append(c, arr1[i])
			i++
		} else {
			c = append(c, arr2[j])
			j++
		}
	}
	c = append(c, arr1[i:]...)
	c = append(c, arr2[j:]...)
	return c
}
