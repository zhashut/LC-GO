package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/12
 * Time: 21:02
 * Description: https://leetcode.cn/problems/spiral-matrix-ii/
 */

func main() {
	generateMatrix(3)
}

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	target := n * n
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for num <= target {
		// 从左到右
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// 从上到下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// 从右到左
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// 从下上到
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
