package dp_maxProduct_152

//乘积最大子数组，要求子数组连续的最大乘积结果
func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	res, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMax, curMin = curMax*nums[i], curMin*nums[i]
		curMax, curMin = max(curMax, curMin, nums[i]), min(curMax, curMin, nums[i])
		if curMax > res {
			res = curMax
		}
	}
	return res
}

func max(x, y, z int) (max int) {
	if x > y {
		max = x
	} else {
		max = y
	}
	if max < z {
		max = z
	}
	return
}

func min(x, y, z int) (min int) {
	if x < y {
		min = x
	} else {
		min = y
	}
	if min > z {
		min = z
	}
	return
}
