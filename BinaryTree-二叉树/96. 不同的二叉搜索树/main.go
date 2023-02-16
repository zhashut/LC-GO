package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/16
 * Time: 0:14
 * Description: 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 */

func main() {
	res := numTrees(3)
	fmt.Println(res)
}

func numTrees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
