package _55__分发饼干

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/25
 * Time: 21:30
 * Description: https://leetcode.cn/problems/assign-cookies/
 */

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	result, index := 0, len(s)-1

	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			result++
			index--
		}
	}

	return result
}
