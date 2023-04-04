package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/4
 * Time: 17:54
 * Description: https://leetcode.cn/problems/distinct-subsequences/
 */

func main() {
	s, t := "rabbbit", "rabbit"
	numDistinct(s, t)
}

// 普通DP：从上到下
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		// dp[i][0] 表示：以i-1为结尾的s可以随便删除元素，出现空字符串的个数。
		// 那么dp[i][0]一定都是1，因为也就是把以i-1为结尾的s，删除所有元素，出现空字符串的个数就是1。
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				// dp[i][j]：以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]。
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
			fmt.Println(dp, i, j)
		}
	}

	return dp[len(s)][len(t)]
}

// 解法二 普通 DP：从下到上
func numDistinct1(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

// 解法一 压缩版 DP
func numDistinct2(s string, t string) int {
	dp := make([]int, len(s)+1)
	for i, curT := range t {
		pre := 0
		for j, curS := range s {
			if i == 0 {
				pre = 1
			}
			newDP := dp[j+1]
			if curT == curS {
				dp[j+1] = dp[j] + pre
			} else {
				dp[j+1] = dp[j]
			}
			pre = newDP
		}
	}
	return dp[len(s)]
}
