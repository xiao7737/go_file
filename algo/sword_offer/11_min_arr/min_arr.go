package _1_min_arr

//二分
//时间O(logN)
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := (left + right) >> 1
		if numbers[mid] > numbers[right] { //最小值在右边
			left = mid + 1
		} else { //最小值在左边
			right = right - 1
		}
	}
	return numbers[left]
}

//暴力，时间O(N)，去遍历，找到下一个数比当前位小的数
