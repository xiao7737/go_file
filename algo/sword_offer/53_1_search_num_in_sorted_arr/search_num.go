package _3_1_search_num_in_sorted_arr

//一个排序数组，统计给定数字在数组中出现的次数

//此种方式，需要单独开map，空间复杂度较高，用二分查找或者前后双指针
func search(nums []int, target int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] != 0 {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	if res, ok := m[target]; ok {
		return res
	} else {
		return 0
	}
}

//利用理解排序好的特点，采用前后双指针的方式
func search2(nums []int, target int) int {
	i, j, res := 0, len(nums)-1, 0
	for i <= j {
		//处理i和target的关系
		if nums[i] > target {
			break
		}
		if nums[i] == target {
			res++
		}
		//处理nuns[]只有一个元素的情况
		if i == j {
			break
		}
		i++

		//处理j和target的关系
		if nums[j] < target {
			break
		}
		if nums[j] == target {
			res++
		}
		j--
	}
	return res
}
