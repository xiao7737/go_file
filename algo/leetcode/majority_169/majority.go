package majority_169

import "sort"

//找到数组中出现次数大于 n/2 的元素，一定有解   [1,1,1,2,3]  返回1

// 排序后直接返回最中间的数
// 时间复杂度NlogN，主要是排序消耗，为快排，空间复杂度O(1)
func Find1(nums []int) int {
	//既然已经超过半数了，那么最中间的数一定是符合要求的数，hhh，有点钻空子的嫌疑，但是是ok的
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 利用map计数
// 时间复杂度O(N)，空间复杂度O(N)，开辟的map使用
func Find2(nums []int) int {
	m := map[int]int{}
	l := len(nums) / 2
	for _, v := range nums {
		// 计数++
		m[v]++
		if m[v] > l {
			return v
		}
	}
	return 0
}
