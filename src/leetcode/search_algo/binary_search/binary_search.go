package main

// 二分查找 logN，要求数组有序
/*func main() {
	var searchKey = 6
	var sortedArr = []int{1, 3, 4, 6, 10, 13}
	result := BinarySearch(sortedArr, searchKey)
	if result >= 0 {
		fmt.Println("Find the index: ", result)
	} else {
		fmt.Println("Not find!")
	}
}*/

func BinarySearch(sortedArr []int, searchKey int) int {
	var low = 0
	var high = len(sortedArr) - 1
	for low <= high {
		var mid = low + (high-low)/2
		if sortedArr[mid] == searchKey {
			return mid
		} else if sortedArr[mid] > searchKey {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 如果low和high特别大，两者之和可能溢出，考虑换成
// mid = low+(high-low)/2
// 更激进的写法，利用位运算代替除法
// mid = low+((high-low)>>1)
