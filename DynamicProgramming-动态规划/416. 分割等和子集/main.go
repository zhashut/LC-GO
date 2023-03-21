package _16__分割等和子集

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/21
 * Time: 18:47
 * Description: https://leetcode.cn/problems/partition-equal-subset-sum/
 */

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	// 如果 nums 的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}

	return dp[target] == target
}
