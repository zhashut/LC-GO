package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/10
 * Time: 20:02
 * Description: https://leetcode.cn/problems/trapping-rain-water/
 */

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(height)
}

// 单调栈
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	sum := 0
	stack := make([]int, 1, len(height)) // 切片模拟单调栈，st存储的是高度数组下标

	for i := 1; i < len(height); i++ {
		top := stack[len(stack)-1]
		if height[i] < height[top] {
			stack = append(stack, i)
		} else if height[i] == height[top] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					l := stack[len(stack)-1]
					h := min(height[l], height[i]) - height[mid]
					w := i - l - 1 // 注意减一，只求中间宽度
					sum += h * w
				}
			}
			stack = append(stack, i)
		}
	}

	return sum
}

// 双指针
func trap2(height []int) int {
	sum := 0
	n := len(height)
	lh := make([]int, n)
	rh := make([]int, n)
	lh[0] = height[0]
	rh[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lh[i] = max(lh[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rh[i] = max(rh[i+1], height[i])
	}
	for i := 1; i < n-1; i++ {
		h := min(rh[i], lh[i]) - height[i]
		if h > 0 {
			sum += h
		}
	}
	return sum
}

// 暴力
func trap3(height []int) int {
	sum := 0

	for i := 0; i < len(height); i++ {
		if i == 0 || i == len(height)-1 {
			continue
		}
		lHeight := height[i]
		rHeight := height[i]
		// 找出左边最高
		for l := i - 1; l >= 0; l-- {
			if height[l] > lHeight {
				lHeight = height[l]
			}
		}
		// 找出右边最高
		for r := i + 1; r < len(height); r++ {
			if height[r] > rHeight {
				rHeight = height[r]
			}
		}
		h := min(lHeight, rHeight) - height[i]
		if h > 0 {
			sum += h
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
