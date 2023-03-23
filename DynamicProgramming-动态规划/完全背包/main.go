package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/23
 * Time: 20:14
 * Description: No Description
 */

func main() {
	weight := []int{1, 3, 4}
	price := []int{15, 20, 30}
	fmt.Println(test_CompletePack1(weight, price, 4))
}

// test_CompletePack1 先遍历物品, 在遍历背包
func test_CompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			fmt.Println(dp)
		}
	}

	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
