package _1_exchange_postion

//将数组中，奇数放在前面，偶数在后面

//双指针，向中间走，交换位置
//时间O(N)
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for nums[i]%2 == 1 {
			//从前到后全是奇数
			if i == j {
				return nums
			}
			i++
		}
		for nums[j]%2 == 0 {
			//从后到前全是偶数
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

// 双指针的另一种写法
func exchange1(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		//找到奇数和偶数，交换
		if nums[i]%2 == 0 && nums[j]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		//从前往后遍历，奇数就i++
		for i < len(nums) && nums[i]%2 != 0 {
			i++
		}
		//从后往前遍历，偶数就j--
		for j >= 0 && nums[j]%2 == 0 {
			j--
		}
	}
	return nums
}
