package main

//从小到大冒泡排序
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] { //从大到小，改为 <
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
