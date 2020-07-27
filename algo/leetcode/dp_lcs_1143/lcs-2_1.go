package dp_lcs_1143

//给定两个字符串 a 和 b，返回这两个字符串的最长公共子串的长度，子串需要连续
// abcd 和 ab 的lcs为2

// dp[i][j] 表示a前i个和b前j个字符最长公共子串的长度
// 构建dp table，如果两个元素相等，dp[i][j] = dp[i-1][j-1] + 1
//   ' a b c d
// ' 0 0 0 0 0
// a 0 1 0 0 0
// b 0 0 2 0 0
func LongestCommonSubsequence2(a string, b string) int {
	if a == "" || b == "" {
		return 0
	}
	dp := make([][]int, len(a)+1)
	maxLen := 0
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// 相等，取左上元素+1
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			} else {
				//只要不相等，直接赋值0
				dp[i][j] = 0
			}
		}
	}
	return maxLen
}
