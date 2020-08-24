package _3_2_miss_num

//找出递增排序数组中缺失的元素，每个元素都唯一，且都是从0开始

//二分法，已经是排序的数组，考虑用二分
//时间O(logN)，空间O(1)
func missNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) >> 1
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

//利用range的k和v相等的关系，时间O(N)，空间O(1)
func missNumber1(nums []int) int {
	res := len(nums)
	for k, v := range nums {
		if k != v {
			return k
		}
	}
	return res
}
