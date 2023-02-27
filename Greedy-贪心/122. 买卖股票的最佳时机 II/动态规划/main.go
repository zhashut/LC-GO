package 动态规划

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/27
 * Time: 19:49
 * Description: No Description
 */

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	// dp[i][0]表示在状态i不持有股票的现金，dp[i][1]为持有股票的现金
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
