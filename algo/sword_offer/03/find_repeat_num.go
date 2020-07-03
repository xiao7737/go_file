package _3

import "sort"

//返回数组中任意一个重复的数字

//利用map
//时间和空间都是O(N)
func FindRepeatNum1(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = true
		}
	}
	return -1
}

//利用排序，排序后相同的元素在一起，判断相邻元素是否相等
//时间O(NlogN)，空间O(1)
func FindRepeatNum2(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
