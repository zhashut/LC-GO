package 链表

import (
	"container/list"
	"sort"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/5
 * Time: 22:14
 * Description: No Description
 */

func reconstructQueue(people [][]int) [][]int {
	// 先将身高从大到小排序，确定最大个子的相对位置
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] //当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0]
	})

	l := list.New() // 创建链表
	for i := 0; i < len(people); i++ {
		position := people[i][1]
		mark := l.PushBack(people[i])
		e := l.Front()
		for position != 0 { // 获取相对位置
			position--
			e = e.Next()
		}
		l.MoveBefore(mark, e) // 移动位置
	}

	result := [][]int{}
	for e := l.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.([]int))
	}

	return result
}
