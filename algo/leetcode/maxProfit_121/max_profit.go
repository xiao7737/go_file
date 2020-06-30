package maxProfit_121

import (
	"math"
)

// 求股票最大收益，只能买卖一次

// 暴力法
// 时间O(N*N)   空间O(1)
func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// 遍历，记录最低点，然后每天都计算最大收益
// 只遍历一次 时间O(N)  空间O(1)
func maxProfit1(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			// 判断是否为最低点
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			// 和之前的最大利润比较，如果比之前大，则更新最大利润
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

// 贪心算法 只要后一天的股价比前一天的高，就卖出
// 更好做法：dp
