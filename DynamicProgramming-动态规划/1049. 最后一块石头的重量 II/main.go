package _049__最后一块石头的重量_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/21
 * Time: 20:15
 * Description: https://leetcode.cn/problems/last-stone-weight-ii/
 */

func lastStoneWeightII(stones []int) int {
	// 题目说最多一共有30个石头，每块石头的最大重量是100，
	//所以极端情况下所有石头相加的重量是 3000，所以我们的大小定义为 3000/2 + 1
	sum := 0
	dp := make([]int, 1501)
	for i := range stones {
		sum += stones[i]
	}

	target := sum / 2

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}

	return sum - 2*dp[target]
}
