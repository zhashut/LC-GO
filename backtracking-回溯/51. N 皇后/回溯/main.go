package 回溯

import "strings"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/24
 * Time: 20:17
 * Description: https://leetcode.cn/problems/n-queens/
 */

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	checkerboard := make([][]string, n, n)

	for i := 0; i < n; i++ {
		checkerboard[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			checkerboard[i][j] = "."
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i, rowStr := range checkerboard {
				tmp[i] = strings.Join(rowStr, "")
			}
			result = append(result, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if isValid(checkerboard, n, row, i) {
				checkerboard[row][i] = "Q"
				backtrack(row + 1)
				checkerboard[row][i] = "."
			}
		}

	}

	backtrack(0)
	return result
}

func isValid(checkerboard [][]string, n int, row int, col int) bool {
	// 判断行
	for i := 0; i < row; i++ {
		if checkerboard[i][col] == "Q" {
			return false
		}
	}
	// 判断左斜线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if checkerboard[i][j] == "Q" {
			return false
		}
	}
	// 判断右斜线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if checkerboard[i][j] == "Q" {
			return false
		}
	}

	return true
}
