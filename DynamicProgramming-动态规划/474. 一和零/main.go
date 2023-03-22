package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/22
 * Time: 20:53
 * Description: https://leetcode.cn/problems/ones-and-zeroes/
 */

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	findMaxForm(strs, m, n)
}

func findMaxForm(strs []string, m int, n int) int {
	// 初始化
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroNum, oneNum := 0, 0
		// 计算0,1 个数
		// 或者直接strings.Count(strs[i],"0")
		for _, v := range str {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(str) - zeroNum
		// 从后往前 遍历背包容量
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				// 推导公式
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
		// fmt.Println(dp)
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
