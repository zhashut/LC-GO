package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/6
 * Time: 20:49
 * Description: https://leetcode.cn/problems/longest-palindromic-subsequence/
 */

func main() {
	s := "bbbab"
	longestPalindromeSubseq(s)
}

func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
			fmt.Println(dp, i, j)
		}
	}

	return dp[0][size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
