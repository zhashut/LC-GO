package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/20
 * Time: 22:09
 * Description: No Description
 */

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test2WeiBagProblem1(weight, value, 4)
}

func test2WeiBagProblem1(weight, value []int, bagweight int) int {
	// 定义dp数组
	dp := make([][]int, len(weight))

	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1)
	}

	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}

	for i := 1; i <= len(weight); i++ {
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[len(weight)-1][bagweight]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
