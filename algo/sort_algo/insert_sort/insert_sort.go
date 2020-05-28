package main

// 插入排序 -- 尾插法  时间复杂度O(n*n)，空间O(1)
// 思路：在未排序的数据中依次取出一个元素，和已经排序好的数组循环比较，交换位置
// 可以参考源码sort的插入排序
func InsertSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}
	for i := 1; i < arrLen; i++ {
		tmp := arr[i] //未排序数据中取的数
		for j := i - 1; j >= 0; j-- {
			if tmp < arr[j] {
				//如果插入的元素比已经排序数组的最后一个元素要小，和已排序数组的元素进行交换，直到tmp>=排序数组的元素
				arr[j], arr[j+1] = tmp, arr[j]
			}
		}
	}
	return arr
}

/*
3 1 4 2
1 3 4 2
1 3 2 4
1 2 3 4
*/
