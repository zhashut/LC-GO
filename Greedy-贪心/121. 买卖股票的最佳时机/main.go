package _21__买卖股票的最佳时机

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/27
 * Time: 20:52
 * Description: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
 */

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return maxProfit
}
