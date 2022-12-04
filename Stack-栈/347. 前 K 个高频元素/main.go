package main

import (
	"container/heap"
	"sort"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/5
 * Time: 0:20
 * Description: No Description
 */

func main() {
	nums, k := []int{1, 1, 1, 2, 2, 3}, 2
	topKFrequent1(nums, k)
}

// 方法一：构建优先队列
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{Key: key, Count: count})
	}

	var result []int
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.Key)
	}
	return result
}

type Item struct {
	Key   int
	Count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Count > pq[j].Count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

//方法二:利用O(logn)排序
func topKFrequent1(nums []int, k int) []int {
	ans := []int{}
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans, func(a, b int) bool {
		return map_num[ans[a]] > map_num[ans[b]]
	})
	return ans[:k]
}
