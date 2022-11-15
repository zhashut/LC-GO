package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/13
 * Time: 21:51
 * Description: No Description
 */

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	spiralOrder(matrix)
}

// 官方
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

//TODO
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, cloums := len(matrix), len(matrix[0])
	left, right := 0, cloums-1
	top, bottom := 0, rows-1
	// 数字的总数量
	res := []int{}
	tarlength := rows * cloums

	// TODO: 这里退出条件有问题
	for len(res) < tarlength {
		// 从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[left][i])
		}
		top++
		// 从上到下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 从右到左
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		// 从下到上
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
