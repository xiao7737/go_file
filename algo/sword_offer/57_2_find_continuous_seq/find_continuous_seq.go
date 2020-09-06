package _7_2_find_continuous_seq

//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列
//输入：target = 9
//输出：[[2,3,4],[4,5]]

//双指针做法
func findContinuousSeq(target int) [][]int {
	left := 1
	right := 2
	var res [][]int
	for left < right {
		//等差求和公式: (首项+末项) * 项数 / 2
		sum := (left + right) * (right - left + 1) / 2
		if sum == target {
			var list []int
			for i := left; i <= right; i++ {
				list = append(list, i)
			}
			res = append(res, list)
			left++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return res
}
