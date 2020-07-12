package _1_exchange_postion

//将数组中，奇数放在前面，偶数在后面

//双指针，向中间走，交换位置
//时间O(N)
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for nums[i]%2 == 1 {
			//全是奇数
			if i == j {
				return nums
			}
			i++
		}
		for nums[j]%2 == 0 {
			//全是偶数
			if j == i {
				return nums
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
