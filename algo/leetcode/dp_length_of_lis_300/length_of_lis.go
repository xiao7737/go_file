package dp_length_of_lis_300

//返回序列中，最长的递增子序列的长度，子序列不用连续
//定义dp[i] 表示第i个元素时，最长的子序列长度
//状态转移：dp[i] = max(dp[i], dp[j+1])
//条件：j < i && nums[i] > nums[j]
//时间O(N*N)，空间O(N)
func LengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		//第一个dp[i]=1
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] { //此时的nums[i]大于前面最大的数，则最长长度+1
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		//更新dp[]里面的最大值
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//利用二分维护一个递增数组，可以将时间复杂度降为NLogN
