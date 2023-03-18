package _3__不同路径_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/18
 * Time: 19:47
 * Description: https://leetcode.cn/problems/unique-paths-ii/
 */

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)

	// 如果起点或终点有障碍物直接返回0
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	// 初始化, 如果是障碍物, 后面的就都是0, 不用循环了
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// // 如果obstacleGrid[i][j]这个点是障碍物, 直接跳过
			if obstacleGrid[i][j] == 1 {
				continue
			}
			// 否则我们需要计算当前点可以到达的路径数
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
