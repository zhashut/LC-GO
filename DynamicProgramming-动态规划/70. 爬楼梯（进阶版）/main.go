package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/24
 * Time: 19:22
 * Description: No Description
 */

func main() {
	n := 2
	climbStairs(n)
}

func climbStairs(n int) int {
	//定义
	dp := make([]int, n+1)
	//初始化
	dp[0] = 1
	// 本题物品只有两个1,2
	m := 2
	for i := 1; i <= n; i++ { // 先遍历背包
		for j := 1; j <= m; j++ { // 再遍历物品
			if i >= j {
				dp[i] += dp[i-j]
			}
			fmt.Println(dp)
		}
	}

	return dp[n]
}
