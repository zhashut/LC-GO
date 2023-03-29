package _88__买卖股票的最佳时机_IV

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/29
 * Time: 20:28
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
 */

// 时间复杂度O(kn) 空间复杂度O(kn)
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][]int, n)
	status := make([]int, (2*k+1)*n)

	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}

	// 奇数天卖出
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 2*k; j += 2 {
			// 奇数天卖出，偶数天选择买或不买
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}

	return dp[n-1][2*k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
