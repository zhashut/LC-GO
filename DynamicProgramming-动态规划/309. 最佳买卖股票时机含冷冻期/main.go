package _09__最佳买卖股票时机含冷冻期

import "math"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/29
 * Time: 21:11
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
 */

// 解题：https://programmercarl.com/0309.%E6%9C%80%E4%BD%B3%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E6%97%B6%E6%9C%BA%E5%90%AB%E5%86%B7%E5%86%BB%E6%9C%9F.html#%E6%80%9D%E8%B7%AF
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[n-1][1], max(dp[n-1][2], dp[n-1][3]))
}

// https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0309.Best-Time-to-Buy-and-Sell-Stock-with-Cooldown/
// 解法二 DP
func maxProfit309(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]
	buy[1] = max(buy[0], -prices[1])
	sell[1] = max(sell[0], buy[0]+prices[1])
	for i := 2; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
	}
	return sell[len(sell)-1]
}

// 解法三 优化辅助空间的 DP
func maxProfit309_1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := []int{-prices[0], max(-prices[0], -prices[1]), math.MinInt64}
	sell := []int{0, max(0, -prices[0]+prices[1]), 0}

	for i := 2; i < len(prices); i++ {
		sell[i%3] = max(sell[(i-1)%3], buy[(i-1)%3]+prices[i])
		buy[i%3] = max(buy[(i-1)%3], sell[(i-2)%3]-prices[i])
	}
	return sell[(len(prices)-1)%3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
