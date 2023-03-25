package _79__完全平方数

import "math"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/25
 * Time: 18:24
 * Description: https://leetcode.cn/problems/perfect-squares/
 */

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 解题二: https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0279.Perfect-Squares/
func numSquares2(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}

// 判断是否为完全平方数
func isPerfectSquare(n int) bool {
	sq := int(math.Floor(math.Sqrt(float64(n))))
	return sq*sq == n
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}
