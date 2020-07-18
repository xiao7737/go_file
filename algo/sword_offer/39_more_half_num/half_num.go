package _9_more_half_num

import "sort"

//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字

//方案1：排序后，中间的那个数即是答案
//时间O(N*logN)		快排 最坏的时间复杂度是O(N*N)，平均O(N*logN)
//空间O(logN)   快排 最坏的空间复杂度是O(N)，平均O(logN)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

//方案2：摩尔投票，对每一个数进行投票，如果该数和下一个数不一样，将投票数-1
//时间O(N)，只需要遍历一次
//空间O(1)
func majorityElement1(nums []int) int {
	x, votes := nums[0], 1
	for i := 1; i < len(nums); i++ {
		//票数归零，开始新一轮计数
		if votes == 0 {
			x = nums[i]
		}

		if nums[i] == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}
