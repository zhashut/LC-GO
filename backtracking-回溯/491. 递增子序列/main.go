package _91__递增子序列

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/22
 * Time: 22:31
 * Description: No Description
 */

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(nums, &res, []int{}, 0)
	return res
}

func backtrack(nums []int, res *[][]int, path []int, start int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	used := make(map[int]bool, len(nums))
	for i := start; i < len(nums); i++ {
		if (used)[nums[i]] { // 去重
			continue
		}

		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			(used)[nums[i]] = true
			backtrack(nums, res, path, i+1)
			path = path[:len(path)-1]
		}
	}
}
