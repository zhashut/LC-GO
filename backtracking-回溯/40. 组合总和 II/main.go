package _0__组合总和_II

import "sort"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/20
 * Time: 19:56
 * Description: No Description
 */

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	used := make([]bool, len(candidates))
	backtrack(candidates, &res, []int{}, target, 0, &used)
	return res
}

func backtrack(candidates []int, result *[][]int, path []int, remain int, startIndex int, used *[]bool) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		if i > 0 && candidates[i] == candidates[i-1] && !(*used)[i-1] {
			continue
		}
		path = append(path, candidates[i])
		(*used)[i] = true
		backtrack(candidates, result, path, remain-candidates[i], i+1, used)
		path = path[:len(path)-1]
		(*used)[i] = false
	}
}
