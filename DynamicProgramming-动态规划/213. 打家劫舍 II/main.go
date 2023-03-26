package _13__打家劫舍_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/26
 * Time: 21:13
 * Description: https://leetcode.cn/problems/house-robber-ii/
 */

// 对于一个数组，成环的话主要有如下三种情况：
// 情况一：考虑不包含首尾元素
// 情况二：考虑包含首元素，不包含尾元素
// 情况三：考虑包含尾元素，不包含首元素
// 注意我这里用的是"考虑"，例如情况三，虽然是考虑包含尾元素，但不一定要选尾部元素！ 对于情况三，取nums[1] 和 nums[3]就是最大的。
// 而情况二 和 情况三 都包含了情况一了，所以只考虑情况二和情况三就可以了.
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[end]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
