package _8__子集

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/21
 * Time: 22:32
 * Description: No Description
 */

func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	backtrack(nums, &res, []int{}, 0)
	return res
}

func backtrack(nums []int, res *[][]int, path []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, res, path, i+1)
		path = path[:len(path)-1]
	}
}
