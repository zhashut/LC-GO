package 贪心

import "math"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/26
 * Time: 18:31
 * Description: https://leetcode.cn/problems/maximum-subarray/
 */

func maxSubArray(nums []int) int {
	result := math.MinInt32
	count := 0

	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count <= 0 {
			count = 0
		}
	}

	return result
}
