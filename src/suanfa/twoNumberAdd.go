package main

import "fmt"

//两数之和，返回符合要求的两个数的index
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func main() {
	s := []int{1, 2, 3, 4, 5}
	res := twoSum(s, 7)
	fmt.Println(res)
}
