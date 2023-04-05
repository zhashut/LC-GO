package _47__回文子串

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/5
 * Time: 20:46
 * Description: https://leetcode.cn/problems/palindromic-substrings/
 */

func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}

// 解题二：
func countSubstrings2(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += countPalindrome(s, i, i)
		res += countPalindrome(s, i, i+1)
	}
	return res
}

func countPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
		res++
	}
	return res
}
