package main

import (
	"fmt"
)

//52. N皇后 II

var ans int
var colRecord map[int]bool //列
var pieRecord map[int]bool //撇 row+col是一样的
var naRecord map[int]bool  //捺 col-row是一样的

func totalNQueens(n int) int {
	ans = 0
	colRecord = make(map[int]bool)
	pieRecord = make(map[int]bool)
	naRecord = make(map[int]bool)
	backtracking(n, 0)
	return ans
}

func backtracking(n int, row int) {
	if n == row {
		ans++
	}
	for col := 0; col < n; col++ {
		if colRecord[col] || pieRecord[col+row] || naRecord[col-row] {
			continue
		}
		colRecord[col] = true
		pieRecord[col+row] = true
		naRecord[col-row] = true
		backtracking(n, row+1)
		colRecord[col] = false
		pieRecord[col+row] = false
		naRecord[col-row] = false
	}
}

func main() {
	fmt.Println(totalNQueens(4))
}
