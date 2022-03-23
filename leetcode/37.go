package main

import "fmt"

func solveSudoku(board [][]byte) {
	if board == nil {
		return
	}
	dfs(board)
}

func dfs(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				continue
			}
			for num := '1'; num <= '9'; num++ {
				if !isValidSudo(board, i, j, byte(num)) {
					continue
				}
				board[i][j] = byte(num)
				if dfs(board) {
					return true
				}
				board[i][j] = '.'
			}
			return false
		}
	}
	return true
}

func isValidSudo(board [][]byte, i, j int, num byte) bool {
	//行
	for row := 0; row < len(board); row++ {
		if board[i][row] == num {
			return false
		}
	}
	//列
	for col := 0; col < len(board); col++ {
		if board[col][j] == num {
			return false
		}
	}
	startCol := i / 3 * 3
	startRow := j / 3 * 3
	for m := startCol; m < startCol+3; m++ {
		for n := startRow; n < startRow+3; n++ {
			if board[m][n] == num {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)

	ret := make([][]string, len(board))
	for i := 0; i < len(board); i++ {
		ret[i] = make([]string, len(board))
		for j := 0; j < len(board); j++ {
			ret[i][j] = string(board[i][j])
		}
	}

	fmt.Println(ret)
}
