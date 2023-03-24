package main

import (
	"fmt"
	"math"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/24
 * Time: 19:50
 * Description: https://leetcode.cn/problems/coin-change/
 */

func main() {
	coins, amount := []int{1, 2, 5}, 11
	coinChange(coins, amount)
}

// 比较喜欢的版本
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
				//fmt.Println(dp, i, j)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 版本一, 先遍历物品,再遍历背包
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				// 打印
				fmt.Println(dp, j, i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// 版本二,先遍历背包,再遍历物品
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	// 遍历背包
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		// 遍历物品
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
				//fmt.Println(dp,j,i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
