package _16__组合总和_III

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/18
 * Time: 19:41
 * Description: No Description
 */

var (
	result [][]int
	path   []int
)

func combinationSum3(k int, n int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, k)

	backTracing(n, k, 0, 1)
	return result
}

func backTracing(targetSum, k, sum, startIndex int) {
	if len(path) == k {
		if targetSum == sum {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
		}
		return
	}

	for i := startIndex; i <= 9; i++ {
		// 剪枝
		if sum+i > targetSum || 9-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		backTracing(targetSum, k, sum+i, i+1)
		path = path[:len(path)-1]
	}
}
