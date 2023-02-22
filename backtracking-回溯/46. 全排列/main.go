package _6__全排列

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/22
 * Time: 23:41
 * Description: No Description
 */

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))

	backtrack(nums, &res, []int{}, &used)
	return res
}

func backtrack(nums []int, res *[][]int, path []int, used *[]bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}
		path = append(path, nums[i])
		(*used)[i] = true
		backtrack(nums, res, path, used)
		(*used)[i] = false
		path = path[:len(path)-1]
	}
}
