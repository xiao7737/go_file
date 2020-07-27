package dp_lcs_1143

// 给定两个字符串 a 和 b，返回这两个字符串的最长公共子串，子串需要连续
// abcd 和 ab 的lcs为ab
// 思路：先求出maxLen，如果table[i][j] == maxLen，则把这个字符放入sli中，
// 在table[i-1][j-1]中继续进行判断，直到table[i][j] < 1

// dp[i][j] 表示a前i个和b前j个字符最长公共子串的长度
// 构建dp table，如果两个元素相等，dp[i][j] = dp[i-1][j-1] + 1
//   ' a b c d
// ' 0 0 0 0 0
// a 0 1 0 0 0
// b 0 0 2 0 0
func LongestCommonSubsequence3(a string, b string) (res string) {
	if a == "" || b == "" {
		return ""
	}
	dp := make([][]int, len(a)+1)
	sli := make([]byte, 0)
	maxLen := 0
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}
	//先求出最长公共子串的长度
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] { // 相等，取左上元素+1
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			} else {
				dp[i][j] = 0 //只要不相等，直接赋值0
			}
		}
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if dp[i][j] == maxLen {
				ii, jj := i, j
				for dp[ii][jj] >= 1 {
					//将该元素转入容器
					sli = append(sli, a[ii-1])
					ii--
					jj--
				}
			}
		}
	}

	//倒序输入sli的元素
	for s := len(sli) - 1; s >= 0; s-- {
		res += string(sli[s])
	}
	return
}
