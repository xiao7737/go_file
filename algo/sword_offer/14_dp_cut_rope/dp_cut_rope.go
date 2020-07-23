package _4_dp_cut_rope

//剪绳子，要求子段相乘的最大结果，要求至少剪一刀
//方案1 dp
//时间O(N*N)，空间O(N)
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	//todo 理解语句，网上搜
	dp := make(map[int]int)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max(res, max(j, dp[j])*max(k, dp[k]))
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]

}

//方案2 贪心
//时间O(1)，空间O(1)
//思路：尽量剪成3，如果长度为4，则剪成2
func cuttingRope2(n int) int {
	if n < 2 {
		return 0
	}
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
