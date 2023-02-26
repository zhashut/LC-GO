package 动态规划

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/26
 * Time: 18:49
 * Description: https://leetcode.cn/problems/maximum-subarray/
 */

func maxSubArray(nums []int) int {
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}
