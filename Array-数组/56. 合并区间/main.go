package main

import "sort"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/2
 * Time: 0:35
 * Description: No Description
 */

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(intervals)
}

//  TODO: 这里测试没问题，力扣的报错: Line 31: Char 19: undefined: merge (solution.go)
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] > intervals[i][1]
	})
	merged := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		L := intervals[i][0]
		R := intervals[i][1]
		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], R)
		}
	}
	return merged
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
