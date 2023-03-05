package main

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/5
 * Time: 21:28
 * Description: https://leetcode.cn/problems/queue-reconstruction-by-height/
 */

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(people)
}

func reconstructQueue(people [][]int) [][]int {
	//  先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // 当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0] // 身高按照由大到小的顺序来排
	})

	// 再按照K进行插入排序，优先插入K小的
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}

	return people
}
