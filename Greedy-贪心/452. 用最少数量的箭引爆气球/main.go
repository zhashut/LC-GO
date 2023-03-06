package main

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/6
 * Time: 20:29
 * Description: No Description
 */

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	findMinArrowShots(points)
}

func findMinArrowShots(points [][]int) int {
	var res int = 1 //弓箭数
	//先按照第一位排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] { //如果前一位的右边界小于后一位的左边界，则一定不重合
			res++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1]) // 更新重叠气球最小右边界,覆盖该位置的值，留到下一步使用
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
