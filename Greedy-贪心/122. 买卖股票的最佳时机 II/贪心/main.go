package 贪心

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/27
 * Time: 19:29
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
 */

// 最大收益来源，必然是每次跌了就买入，涨到顶峰的时候就抛出。
// 只要有涨峰就开始计算赚的钱，连续涨可以用两两相减累加来计算，
// 两两相减累加，相当于涨到波峰的最大值减去谷底的值。这一点看通以后，题目非常简单.
func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
