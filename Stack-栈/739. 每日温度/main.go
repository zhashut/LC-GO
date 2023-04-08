package _39__每日温度

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/8
 * Time: 21:33
 * Description: https://leetcode.cn/problems/daily-temperatures/
 */

// 单调栈法(未精简版本)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 初始化栈顶元素为第一个下标索引0
	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}

// 单调递减栈(精简版)
func dailyTemperatures2(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
