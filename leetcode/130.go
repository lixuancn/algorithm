package main

import (
	"container/list"
	"fmt"
)

//130. 被围绕的区域

func solve(board [][]byte) {
	solve_dfs(board)
	solve_bfs(board)
}

//靠边的O不能被染色，那么就找到靠边的O，然后看这个O和那些O相连，把他们做个标记，最后把非标记的O染成X
func solve_dfs(board [][]byte) {
	if len(board) == 0 || len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		solve_dfs_handler(board, i, 0)
		solve_dfs_handler(board, i, len(board[0])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		solve_dfs_handler(board, 0, j)
		solve_dfs_handler(board, len(board)-1, j)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solve_dfs_handler(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	solve_dfs_handler(board, i+1, j)
	solve_dfs_handler(board, i, j+1)
	solve_dfs_handler(board, i-1, j)
	solve_dfs_handler(board, i, j-1)
}

func solve_bfs(board [][]byte) {
	if len(board) == 0 || len(board) == 0 {
		return
	}
	queue := list.New()
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			queue.PushBack([]int{i, 0})
			board[i][0] = 'A'
		}
		if board[i][len(board[0])-1] == 'O' {
			queue.PushBack([]int{i, len(board[0]) - 1})
			board[i][len(board[0])-1] = 'A'
		}
	}
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			queue.PushBack([]int{0, j})
			board[0][j] = 'A'
		}
		if board[len(board)-1][j] == 'O' {
			queue.PushBack([]int{len(board) - 1, j})
			board[len(board)-1][j] = 'A'
		}
	}
	for queue.Len() > 0 {
		coordinate := queue.Remove(queue.Front()).([]int)
		for _, pos := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			newI, newJ := coordinate[0]+pos[0], coordinate[1]+pos[1]
			if newI < 0 || newI >= len(board) || newJ < 0 || newJ >= len(board[0]) || board[newI][newJ] != 'O' {
				continue
			}
			queue.PushBack([]int{newI, newJ})
			board[newI][newJ] = 'A'
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	solve(board)
	fmt.Println(board)
}
