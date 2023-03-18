package _2__不同路径

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/18
 * Time: 19:15
 * Description: https://leetcode.cn/problems/unique-paths/
 */

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
