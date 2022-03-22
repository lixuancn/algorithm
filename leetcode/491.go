package main

import (
	"fmt"
)

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	row := make([]int, 0)
	backtracking(nums, 0, row, &result)
	return result
}

func backtracking(nums []int, start int, row []int, result *[][]int) {
	if len(row) >= 2 {
		tmp := make([]int, len(row))
		copy(tmp, row)
		*result = append(*result, tmp)
	}
	//本层的去重，去重和和90、40不相同
	//题目是-100 ~ 100
	record := make([]int, 201)
	for i := start; i < len(nums); i++ {
		//1、如果当前元素小于子集的最后一个元素，则跳过
		//2、当前元素在本层已经出现过，所以跳过后继续寻找
		if (len(row) > 0 && nums[i] < row[len(row)-1]) || record[nums[i]+100] == 1 {
			continue
		}
		row = append(row, nums[i])
		record[nums[i]+100] = 1
		backtracking(nums, i+1, row, result)
		row = row[:len(row)-1]
	}
}

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
