package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/10/27
 * Time: 23:56
 * Description: No Description
 */

func maxArea(height []int) int {
	multiple, start, end := 0, 0, len(height)-1

	for start < end {
		width := end - start
		high := 0

		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		multiple = max(multiple, width*high)
	}
	return multiple
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
