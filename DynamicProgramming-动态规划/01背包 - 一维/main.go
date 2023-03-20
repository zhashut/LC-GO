package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/20
 * Time: 22:47
 * Description: No Description
 */

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test1WeiBagProblem(weight, value, 4)
}

func test1WeiBagProblem(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)

	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	fmt.Println(dp)
	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
