package main

// 二分查找 logN，要求数组有序
// 可优化点：
// 如果low和high特别大，两者之和可能溢出，考虑换成
// mid = low+(high-low)/2
// 更激进的写法，利用位运算代替除法
// mid = low+((high-low)>>1)

func BinarySearch(sortedArr []int, searchKey int) int {
	var low = 0
	var high = len(sortedArr) - 1
	for low <= high {
		var mid = low + (high-low)/2
		if sortedArr[mid] == searchKey {
			return mid
		} else if sortedArr[mid] > searchKey {
			high = mid - 1 //此处不写 high = mid，避免出现死循环：low=high，而sortedArr[mid]!=searchKey，导致一直循环
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 变种1：查找第一个searchKey出现的元素（数组有重复元素）
func BinarySearch1(sortedArr []int, searchKey int) int {
	var low = 0
	var high = len(sortedArr) - 1
	for low <= high {
		var mid = low + ((high - low) >> 1)
		if sortedArr[mid] > searchKey {
			high = mid - 1
		} else if sortedArr[mid] < searchKey {
			low = mid + 1
		} else {
			// 如果mid为0，肯定是第一个了；如果该mid的前一个数不是search，那么该mid就是第一个出现的searchKey
			if (mid == 0) || (sortedArr[mid-1] != searchKey) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 变种2：查找第一个大于searchKey的元素
func BinarySearch2(sortedArr []int, searchKey int) int {
	var low = 0
	var high = len(sortedArr) - 1
	for low <= high {
		var mid = low + ((high - low) >> 1)
		if sortedArr[mid] >= searchKey {
			// mid的前一个小于searchKey，而当前mid大于searchKey
			if (mid == 0) || (sortedArr[mid-1] < searchKey) {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}
