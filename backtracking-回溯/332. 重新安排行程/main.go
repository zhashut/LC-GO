package _32__重新安排行程

import (
	"sort"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/23
 * Time: 22:53
 * Description: https://leetcode.cn/problems/reconstruct-itinerary/
 */

type pair struct {
	target  string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findItinerary(tickets [][]string) []string {
	result := []string{}
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	// 初始化 targets 数组
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}

	// 字母序排序
	for s := range targets {
		sort.Sort(targets[s])
	}

	result = append(result, "JFK")
	var backtracking func() bool
	backtracking = func() bool {
		if len(result) == len(tickets)+1 {
			return true
		}

		// 取出起飞航班对应的目的地
		for _, pair := range targets[result[len(result)-1]] {
			// 如果该航班没有飞过
			if !pair.visited {
				result = append(result, pair.target)
				pair.visited = true
				if backtracking() {
					return true
				}
				result = result[:len(result)-1]
				pair.visited = false
			}
		}
		return false
	}

	backtracking()

	return result
}

func findItinerary2(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
