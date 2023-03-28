package _23__买卖股票的最佳时机_III

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/28
 * Time: 20:12
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
 */

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[n-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
