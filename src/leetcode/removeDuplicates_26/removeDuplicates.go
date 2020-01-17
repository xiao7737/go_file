package removeDuplicates_26

// 题目：给定一个已排序数组，原地去除重复的元素，使得每个元素只出现一次，返回新数组的新长度，要求空间O(1)
// arr = [1,2,2,3,4]   返回4
// 思路：双指针，不用真正的删除，只返回长度就行

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	// 两个指针同时走，一个快，一个慢
	// j是快指针，i是慢指针，每次循环，j都++，而i只有在不同的情况下才++，直到j遍历完数组
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j] //如果不一样，则将新的元素赋值给慢指针元素
		}
	}
	return i + 1
}
