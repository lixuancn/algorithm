package main

import "fmt"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	row := make([]int, 0)
	backtracking(nums, row, &result, 0)
	return result
}

func backtracking(nums, row []int, result *[][]int, start int) {
	tmp := make([]int, len(row))
	fmt.Println(row)
	copy(tmp, row)
	*result = append(*result, tmp)
	if start == len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		row = append(row, nums[i])
		backtracking(nums, row, result, i+1)
		row = row[:len(row)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
