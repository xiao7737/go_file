package threeNumberAdd_15

import (
	"sort"
)

// 题目：对于无序数组，找到满足 i + j + k = target 的元素组，且i j k三元组每一组不相等

//方案1：三次循环，O(N*N*N)
func ThreeAdd(nums []int, target int) []int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if ((i != j) && (i != k) && (j != k)) && (nums[i]+nums[j]+nums[k] == target) {
					return []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}
	return nil
}

//方案2：枚举a和b，利用target-(a+b)来找c，两数之和的TwoNum4的解法
//O(N*N)，空间O(N)，会利用一个map来存N个数
func ThreeAdd1(nums []int, target int) []int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	//todo
	return nil
}

//方案3：对数组排序，枚举a，利用指针对撞来找b和c，两数之和的TwoSum2的解法
//O(N*N)，空间O(1)
//找到一组满足的元素
func ThreeAdd2(nums []int, target int) (res []int) {
	length := len(nums)
	if length < 3 {
		return nil
	}

	//对数组排序
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		j, k := i+1, length-1
		for j < k {
			sumJK := nums[j] + nums[k]
			if sumJK > target-nums[i] {
				k--
			} else if sumJK < target-nums[i] {
				i++
			} else {
				return []int{nums[i], nums[j], nums[k]}
			}
		}
	}
	return nil
}

// 找出多组满足的元素
// 可以考虑用map来去重
// todo 还有问题bug   136的组合
func ThreeAdd3(nums []int, target int) (res [][]int) {
	length := len(nums)
	if length < 3 {
		return nil
	}
	//对数组排序
	sort.Ints(nums)

	for i := 0; i < length-2; i++ {
		//i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > target {
			break
		}

		j, k := i+1, length-1
		for j < k {
			sumJK := nums[j] + nums[k]
			if sumJK > target-nums[i] {
				k--
				continue
			} else if sumJK < target-nums[i] {
				i++
				continue
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				//j去重
				for {
					if j < k && nums[j] == nums[j+1] {
						j++
					} else {
						break
					}
				}
				//k去重
				for {
					if j < k && nums[k] == nums[k-1] {
						k--
					} else {
						break
					}
				}
				j++
				k--
			}
		}
	}
	return res
}
