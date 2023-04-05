package _2__编辑距离

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/5
 * Time: 18:11
 * Description: https://leetcode.cn/problems/edit-distance/
 */

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	// 初始化
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			//fmt.Println(dp, i, j)
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
