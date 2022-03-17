package main

import (
	"fmt"
	"sort"
)

//回溯 40.组合总和II
var result [][]int
var row []int
var used map[int]bool

func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	row = make([]int, 0)
	used = make(map[int]bool)
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
		//这是排重，抽象成树的话，就是同一层的元素不能相同
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		//if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
		//	continue
		//}
		//也可以用start来判断
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		sum += candidates[i]
		row = append(row, candidates[i])
		used[i] = true
		backtracking(candidates, target, i+1, sum)
		row = row[:len(row)-1]
		sum -= candidates[i]
		used[i] = false
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
