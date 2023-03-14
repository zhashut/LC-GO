package _7__全排列_II

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/14
 * Time: 0:54
 * Description: https://leetcode.cn/problems/permutations-ii/
 */

var (
	res  [][]int
	path []int
	st   []bool // state的缩写
)

func permuteUnique(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, cur int) {
	if cur == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] && !st[i-1] { // 去重，用st来判别是深度还是广度
			continue
		}
		if !st[i] {
			path = append(path, nums[i])
			st[i] = true
			dfs(nums, cur+1)
			st[i] = false
			path = path[:len(path)-1]
		}
	}
}
