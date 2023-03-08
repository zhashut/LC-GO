package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/7
 * Time: 22:32
 * Description: https://leetcode.cn/problems/partition-labels/
 */

func main() {
	s := "ababcbacadefegdehijhklij"
	partitionLabels(s)
}

func partitionLabels(s string) []int {
	var res []int
	var marks [26]int
	size, left, right := len(s), 0, 0

	for i := 0; i < size; i++ {
		marks[s[i]-'a'] = i
	}

	for i := 0; i < size; i++ {
		right = max(right, marks[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
