package main

import (
	"math"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/22
 * Time: 20:07
 * Description: https://leetcode.cn/problems/target-sum/
 */

func main() {
	nums, target := []int{1, 1, 1, 1, 1}, 3
	findTargetSumWays(nums, target)
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	// 绝对值
	if abs(target) > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bag+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			// 推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp) // 打印，看推导过程，没用的时候就去掉
		}
	}

	return dp[bag]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
