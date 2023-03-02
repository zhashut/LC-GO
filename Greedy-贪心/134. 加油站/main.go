package _34__加油站

import "math"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/1
 * Time: 20:56
 * Description: https://leetcode.cn/problems/gas-station/submissions/
 */

// 解法一：全局最优出发，没有局部最优，归类为属于贪心的一半吧
func canCompleteCircuit(gas []int, cost []int) int {
	curSum, min := 0, math.MaxInt32

	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		curSum += rest
		if curSum < min {
			min = curSum
		}
	}

	// 如果总和为负数，则说明走不完一周
	if curSum < 0 {
		return -1
	}
	// 如果 min 大于或等于0，则起点为0
	if min >= 0 {
		return 0
	}
	// 如果 min 小于0，则倒序遍历，能把 min 抹平的则为起点
	for i := len(gas) - 1; i >= 0; i-- {
		rest := gas[i] - cost[i]
		min += rest
		if min >= 0 {
			return i
		}
	}

	return -1
}

// 解法二：从局部最优退出全局最优
func canCompleteCircuit2(gas []int, cost []int) int {
	curSum, totalSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}
