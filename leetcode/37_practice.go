package main

import "fmt"

//37. 解数独
//把37.go中的循环判断改为hash表，把全board循环改为只循环空格

var rowRecord, colRecord [9][9]bool
var block [3][3][9]bool
var spaces [][2]int

func solveSudoku(board [][]byte) {
	if board == nil {
		return
	}
	rowRecord, colRecord = [9][9]bool{}, [9][9]bool{}
	block = [3][3][9]bool{}
	spaces = [][2]int{}
	//先找到所有的空点，并初始化已有数的记录用来排重
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := int(board[i][j] - '1')
				rowRecord[i][digit] = true
				colRecord[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}
	backtracking(board, 0)
}

func backtracking(board [][]byte, pos int) bool {
	if pos == len(spaces) {
		return true
	}
	i, j := spaces[pos][0], spaces[pos][1]
	for num := byte(0); num < 9; num++ {
		if rowRecord[i][num] || colRecord[j][num] || block[i/3][j/3][num] {
			continue
		}
		rowRecord[i][num] = true
		colRecord[j][num] = true
		block[i/3][j/3][num] = true
		board[i][j] = num + '1' //board不需要回溯，因为是按照空格spaces来循环的
		if backtracking(board, pos+1) {
			return true
		}
		rowRecord[i][num] = false
		colRecord[j][num] = false
		block[i/3][j/3][num] = false
	}
	return false
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
