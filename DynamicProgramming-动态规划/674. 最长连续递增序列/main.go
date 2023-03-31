package _74__最长连续递增序列

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/31
 * Time: 12:46
 * Description: https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
 */

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}
