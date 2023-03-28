package _22__买卖股票的最佳时机_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/28
 * Time: 19:43
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
 */

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0] // 持有股票的当前现金
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[n-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
