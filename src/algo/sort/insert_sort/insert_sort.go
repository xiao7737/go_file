package main

// 插入排序 -- 尾插法  时间复杂度O(n*n)，空间O(1)
// 思路：在未排序的数据中取出，插入到已经排序好的数组中
// 可以参考源码sort的插入排序
func InsertSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}
	for i := 1; i < arrLen; i++ {
		tmp := arr[i] //未排序数据中取的数
		for j := i - 1; j >= 0; j-- {
			if tmp >= arr[j] {
				//如果插入的元素tmp>=已经排序数组的最后一个元素,那么直接不做交换，也不做对比
				break
			} else {
				//如果插入的元素比已经排序数组的最后一个元素要小，逐渐和已排序数组的元素进行交换，直到tmp>=排序数组的元素
				arr[j+1], arr[j] = arr[j], tmp
			}
		}
	}
	return arr
}
