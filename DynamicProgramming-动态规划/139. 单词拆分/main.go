package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/25
 * Time: 18:51
 * Description: https://leetcode.cn/problems/word-break/
 */

func main() {
	s, wordDict := "leetcode", []string{"leet", "code"}
	wordBreak(s, wordDict)
}

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
		fmt.Println(dp)
	}

	return dp[len(s)]
}

// 解题二： 转化为 求装满背包s的前几位字符的方式有几种
func wordBreak2(s string, wordDict []string) bool {
	// 装满背包s的前几位字符的方式有几种
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i <= len(s); i++ { // 背包
		for j := 0; j < len(wordDict); j++ { // 物品
			if i >= len(wordDict[j]) && wordDict[j] == s[i-len(wordDict[j]):i] {
				dp[i] += dp[i-len(wordDict[j])]
			}
		}
	}

	return dp[len(s)] > 0
}
