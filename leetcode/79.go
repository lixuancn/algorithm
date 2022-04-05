package main

import "fmt"

//79. 单词搜索

var used [][]bool

func exist(board [][]byte, word string) bool {
	used = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			used[i][j] = true
			if backtracking(board, word, i, j, []byte{board[i][j]}) {
				return true
			}
			used[i][j] = false
		}
	}
	return false
}

func backtracking(board [][]byte, word string, i, j int, ret []byte) bool {
	if ret[len(ret)-1] != word[len(ret)-1] {
		return false
	}
	if string(ret) == word {
		return true
	}
	if len(ret) >= len(word) {
		return false
	}
	for _, pos := range [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newI >= len(board) || newJ < 0 || newJ >= len(board[0]) {
			continue
		}
		if used[newI][newJ] {
			continue
		}
		used[newI][newJ] = true
		ret = append(ret, board[newI][newJ])
		if backtracking(board, word, newI, newJ, ret) {
			return true
		}
		ret = ret[:len(ret)-1]
		used[newI][newJ] = false
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}
