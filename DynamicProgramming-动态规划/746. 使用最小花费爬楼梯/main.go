package _46__使用最小花费爬楼梯

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/17
 * Time: 22:37
 * Description: https://leetcode.cn/problems/min-cost-climbing-stairs/
 */

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)

	// 初始化
	dp[0], dp[1] = 0, 0
	// 遍历
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
