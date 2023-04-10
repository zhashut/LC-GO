package _4__柱状图中最大的矩形

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/10
 * Time: 22:04
 * Description: https://leetcode.cn/problems/largest-rectangle-in-histogram/
 */

// 单调栈
// https://programmercarl.com/0084.%E6%9F%B1%E7%8A%B6%E5%9B%BE%E4%B8%AD%E6%9C%80%E5%A4%A7%E7%9A%84%E7%9F%A9%E5%BD%A2.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
func largestRectangleArea(heights []int) int {
	// 初始化栈
	max, stack := 0, []int{0}
	// 数组头部加入0
	heights = append([]int{0}, heights...)
	// 数组尾部加入0
	heights = append(heights, 0)

	for i := 1; i < len(heights); i++ {
		for heights[stack[len(stack)-1]] > heights[i] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				left := stack[len(stack)-1]
				tmp := heights[mid] * (i - left - 1)
				if tmp > max {
					max = tmp
				}
			}
		}
		stack = append(stack, i)
	}

	return max
}

// 单调栈 + 常数优化
// 链接：https://leetcode.cn/problems/largest-rectangle-in-histogram/solution/zhu-zhuang-tu-zhong-zui-da-de-ju-xing-by-leetcode-/
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
