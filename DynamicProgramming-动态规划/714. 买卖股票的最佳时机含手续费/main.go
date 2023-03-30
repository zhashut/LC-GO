package _14__买卖股票的最佳时机含手续费

import "math"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/30
 * Time: 20:05
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 */

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(prices)-1][1]
}

// 解法一 模拟 DP
func maxProfit714(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
	}
	return sell[len(sell)-1]
}

// 解法二 优化辅助空间的 DP
func maxProfit714_1(prices []int, fee int) int {
	sell, buy := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
