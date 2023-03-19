package _6__不同的二叉搜索树

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/19
 * Time: 18:50
 * Description: https://leetcode.cn/problems/unique-binary-search-trees/
 */

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j] // 左子树 * 右子树
		}
	}

	return dp[n]
}
