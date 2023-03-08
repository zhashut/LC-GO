package _6__合并区间

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/8
 * Time: 21:07
 * Description: https://leetcode.cn/problems/merge-intervals/
 */

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0, len(intervals))
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			result = append(result, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	result = append(result, []int{left, right}) // 将最后一个区间放入

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
