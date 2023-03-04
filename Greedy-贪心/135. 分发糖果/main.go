package _35__分发糖果

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/3
 * Time: 23:58
 * Description: No Description
 */

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/4
 * Time: 21:46
 * Description: https://leetcode.cn/problems/candy/
 */

func candy(ratings []int) int {
	candyVec, sum := make([]int, len(ratings)), 0

	// 初始化(每个人至少一个糖果)
	for i := 0; i < len(ratings); i++ {
		candyVec[i] = 1
	}
	// 1.先从左到右，当右边的大于左边的就加1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyVec[i] = candyVec[i-1] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边的就右边加1，但要花费糖果最少，所以需要做下判断
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyVec[i] = max(candyVec[i], candyVec[i+1]+1)
		}
	}
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		sum += candyVec[i]
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
