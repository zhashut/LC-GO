package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/23
 * Time: 20:59
 * Description: https://leetcode.cn/problems/combination-sum-iv/
 */

func main() {
	nums, target := []int{1, 2, 3}, 4
	combinationSum4(nums, target)
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
			fmt.Println(dp)
		}
	}

	return dp[target]
}
