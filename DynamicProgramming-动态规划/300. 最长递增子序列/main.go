package _00__最长递增子序列

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/30
 * Time: 20:36
 * Description: https://leetcode.cn/problems/longest-increasing-subsequence/
 */

func lengthOfLIS(nums []int) int {
	// dp数组的定义 dp[i]表示取第i个元素的时候，表示子序列的长度，其中包括 nums[i] 这个元素
	dp := make([]int, len(nums))
	// 初始化，所有的元素都应该初始化为1
	for i := range dp {
		dp[i] = 1
	}

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

// https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0300.Longest-Increasing-Subsequence/
// 解法二 O(n log n) DP
func lengthOfLIS1(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
