package twoNumberAdd

//要求：输入数组，目标值，返回符合数组中：两个数相加等于目标值的元素的index
//两数之和，返回符合要求的两个数的index，且两个index不重复，只返回符合要求的一组数

//暴力解法，O(N*N)
//分别对两个数枚举
func TwoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//指针对撞，O(N)
//只适合nums有序的情况下，只需要枚举一个数
//寻找nums[i]+nums[j]=target，从两头分别向中间找
func TwoSum2(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	sum := nums[i] + nums[j]
	for sum != target {
		if sum < target {
			i++
		} else {
			j--
		}
		sum = nums[i] + nums[j]
	}
	return []int{i, j}
}

//二分查找 O(NlogN)
//适合数组有序
func TwoNum3(nums []int, target int) []int {
	length := len(nums)
	l2 := length
	for i := 0; i < length; i++ {
		v := target - nums[i]
		for j := i + 1; j < l2; {
			mid := j + (l2-j)/2 //避免内存溢出
			if v == nums[mid] {
				return []int{i + 1, mid + 1}
			} else if v < nums[mid] { // v在左边
				l2 = mid
			} else { // v在右边
				j = mid + 1
			}
		}
	}
	return []int{}
}

//利用map，也可以实现O(N)，枚举一个数即可
//nums可以无序
/*
思路：
将x装入map
判断 target-x的结果在不在map中*/
