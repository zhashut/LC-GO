package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/23
 * Time: 20:22
 * Description: https://leetcode.cn/problems/coin-change-ii/
 */

func main() {
	amount, coins := 5, []int{1, 2, 5}
	change(amount, coins)
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
			//fmt.Println(dp)
		}
	}

	return dp[amount]
}
