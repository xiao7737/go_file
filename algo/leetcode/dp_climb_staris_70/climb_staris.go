package dp_climb_staris_70

// 爬楼梯，每次只能走1阶或者2阶，给定n阶的楼梯，求总的走法
// 采用动态规划进行推导
// 时间O(N)，空间O(n)常数级
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	mem := make([]int, n)
	mem[0] = 1
	mem[1] = 2
	for i := 2; i < n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n-1]
}
