package _0__子集_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/14
 * Time: 0:31
 * Description: No Description
 */

var (
	res  [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)

	backtrack(nums, 0)
	return res
}

func backtrack(nums []int, start int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}

	used := make(map[int]bool, len(nums))
	for i := start; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			backtrack(nums, i+1)
			path = path[:len(path)-1]
		}
	}
}
