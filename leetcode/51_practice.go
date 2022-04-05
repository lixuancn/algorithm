package main

import (
	"fmt"
)

//51. N 皇后

var result [][]int
var colRecord map[int]bool //列
var pieRecord map[int]bool //撇 row+col是一样的
var naRecord map[int]bool  //捺 col-row是一样的

func solveNQueens(n int) [][]string {
	result = make([][]int, 0)
	ret := make([]int, 0)
	colRecord = make(map[int]bool)
	pieRecord = make(map[int]bool)
	naRecord = make(map[int]bool)
	backtracking(n, 0, ret)
	ansList := make([][]string, len(result))
	for i := 0; i < len(result); i++ {
		ans := make([]string, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if k == result[i][j] {
					ans[k] += "Q"
				} else {
					ans[k] += "."
				}
			}
		}
		ansList[i] = ans
	}
	return ansList
}

func backtracking(n int, row int, ret []int) {
	if row == n {
		fmt.Println(ret)
		result = append(result, append([]int(nil), ret...))
		return
	}
	for col := 0; col < n; col++ {
		if colRecord[col] || pieRecord[row+col] || naRecord[col-row] {
			continue
		}
		ret = append(ret, col)
		colRecord[col] = true
		pieRecord[row+col] = true
		naRecord[col-row] = true
		backtracking(n, row+1, ret)
		colRecord[col] = false
		pieRecord[row+col] = false
		naRecord[col-row] = false
		ret = ret[:len(ret)-1]
	}
}

func main() {
	fmt.Println(solveNQueens(4))
}
