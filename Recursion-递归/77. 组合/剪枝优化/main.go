package 剪枝优化

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/18
 * Time: 19:30
 * Description: No Description
 */

var (
	result [][]int
	path   []int
)

func combine(n int, k int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, k)

	backTracing(n, k, 1)
	return result
}

func backTracing(n, k, startIndex int) {
	if len(path) == k {
		temp := make([]int, 0)
		copy(temp, path)
		result = append(result, temp)
		return
	}

	for i := startIndex; i < n-(k-len(path))+1; i++ {
		path = append(path, i)
		backTracing(n, k, i+1)
		path = path[:len(path)-1]
	}
}
