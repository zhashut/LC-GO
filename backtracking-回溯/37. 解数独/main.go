package _7__解数独

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/24
 * Time: 22:03
 * Description: https://leetcode.cn/problems/sudoku-solver/
 */

func solveSudoku(board [][]byte) {
	var backtrack func([][]byte) bool

	backtrack = func(board [][]byte) bool {
		for i := 0; i < len(board); i++ { // 行
			for j := 0; j < len(board[0]); j++ { // 列
				if board[i][j] != '.' { // 判断此位置是否适合填数字
					continue
				}
				for k := '1'; k <= '9'; k++ { // 尝试填1-9
					if isValid(i, j, board, byte(k)) {
						board[i][j] = byte(k)
						if backtrack(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}

	backtrack(board)
}

func isValid(row int, col int, board [][]byte, k byte) bool {
	// 同一行每一列
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] == k {
			return false
		}
	}
	// 同一列每一行
	for i := 0; i < len(board); i++ {
		if board[i][col] == k {
			return false
		}
	}
	// 九宫格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}

	return true
}
