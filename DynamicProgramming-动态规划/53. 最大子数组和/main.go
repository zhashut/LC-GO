package _3__最大子数组和

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/2
 * Time: 20:50
 * Description: https://leetcode.cn/problems/maximum-subarray/
 */

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	maxVal := 0
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxVal {
			maxVal = dp[i]
		}
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
