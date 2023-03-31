package _18__最长重复子数组

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/31
 * Time: 13:01
 * Description: https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
 */

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}

	return result
}

// 滚动数组
func findLength2(nums1 []int, nums2 []int) int {
	n, m, res := len(nums1), len(nums2), 0
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := m; j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0 // 注意这里不相等要赋值为0，供下一层使用
			}
			res = max(res, dp[j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
