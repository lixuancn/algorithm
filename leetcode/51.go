package main

import (
	"fmt"
	"strconv"
)

//51. N 皇后

var usedCol map[int]bool
var usedPie map[int]bool
var usedNa map[int]bool
var result [][]string
var oneAns []string

func solveNQueens(n int) [][]string {
	result = [][]string{}
	oneAns = []string{}
	usedCol = map[int]bool{}
	usedPie = map[int]bool{}
	usedNa = map[int]bool{}
	dfs(n, 0, oneAns)

	ret := make([][]string, len(result))
	for i := 0; i < len(result); i++ {
		ret[i] = make([]string, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if strconv.Itoa(k) == result[i][j] {
					ret[i][j] += "Q"
				} else {
					ret[i][j] += "."
				}
			}
		}
	}
	return ret
}

func dfs(n int, row int, oneAns []string) {
	if row == n {
		tmp := make([]string, len(oneAns))
		copy(tmp, oneAns)
		result = append(result, tmp)
		return
	}
	for col := 0; col < n; col++ {
		if usedCol[col] || usedPie[col+row] || usedNa[col-row] {
			continue
		}
		usedCol[col] = true
		usedPie[col+row] = true
		usedNa[col-row] = true
		oneAns = append(oneAns, strconv.Itoa(col))
		dfs(n, row+1, oneAns)
		oneAns = oneAns[:len(oneAns)-1]
		usedCol[col] = false
		usedPie[col+row] = false
		usedNa[col-row] = false
	}
}

func main() {
	fmt.Println(solveNQueens(4))
}
