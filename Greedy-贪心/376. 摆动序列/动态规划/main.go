package 动态规划

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/25
 * Time: 22:57
 * Description: https://leetcode.cn/problems/wiggle-subsequence/
 */

/*
https://programmercarl.com/0376.%E6%91%86%E5%8A%A8%E5%BA%8F%E5%88%97.html#%E6%80%9D%E8%B7%AF1-%E8%B4%AA%E5%BF%83%E8%A7%A3%E6%B3%95
时间复杂度：O(n^2)
空间复杂度：O(n)
*/
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([][2]int, n)
	// i 0 作为波峰的最大长度
	// i 1 作为波谷的最大长度
	dp[0][0], dp[0][1] = 1, 1

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] { // nums[i]作为波峰
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
			if nums[i] < nums[j] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
			if nums[i] == nums[j] { // 添加一种情况，nums[i]为相等
				dp[i][0] = max(dp[i][0], dp[j][0]) // 波峰
				dp[i][1] = max(dp[i][1], dp[j][1]) // 波谷
			}
		}
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
