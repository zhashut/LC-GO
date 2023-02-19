package _7__组合

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/17
 * Time: 1:37
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
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backTracing(n, k, i+1)
		path = path[:len(path)-1]
	}
}
