package 贪心

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/27
 * Time: 19:29
 * Description: No Description
 */

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
