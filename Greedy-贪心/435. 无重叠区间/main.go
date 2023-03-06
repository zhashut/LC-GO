package main

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/6
 * Time: 21:28
 * Description: No Description
 */

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	eraseOverlapIntervals(intervals)
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if end <= intervals[i][0] {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}
