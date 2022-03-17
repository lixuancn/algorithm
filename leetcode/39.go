package main

import (
	"fmt"
	"sort"
)

//回溯 39. 组合总和
var result [][]int
var row []int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	row = make([]int, 0)
	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0)
	return result
}

func backtracking(candidates []int, target, start, sum int) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(row))
		copy(tmp, row)
		result = append(result, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		sum += candidates[i]
		row = append(row, candidates[i])
		backtracking(candidates, target, i, sum)
		row = row[:len(row)-1]
		sum -= candidates[i]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
