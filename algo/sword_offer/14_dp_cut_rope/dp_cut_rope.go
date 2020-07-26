package _4_dp_cut_rope

import "math"

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

	dp := make(map[int]int)
	dp[1] = 0
	dp[2] = 1
	dp[3] = 2

	for i := 4; i <= n; i++ {
		j, k := 1, i-1
		res := 0
		for j = 1; j <= n/2; i++ {
			//理解为剪成两段相乘，但是两段分别要为最大，每一段又由更小的两段相乘得出，形成重复子问题
			res = max(res, max(j, dp[j])*max(k, dp[k]))
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

//方案2 贪心
//时间O(1)，空间O(1)
//思路：尽量剪成3，如果长度为4，直接是2*2=4
func cuttingRope2(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	parts := n / 3 // 能分成几份3
	another := n % 3
	var result float64

	switch another {
	case 2:
		result = math.Pow(3, float64(parts))
		result *= 2
	case 1:
		result = math.Pow(3, float64(parts-1))
		result *= 4
	default:
		result = math.Pow(3, float64(parts))
	}
	return int(result)
}
