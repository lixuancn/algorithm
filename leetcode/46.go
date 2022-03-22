package main

import "fmt"

func permute(nums []int) [][]int {
	row := make([]int, 0)
	result := make([][]int, 0)
	record := make([]uint8, len(nums))
	backtracking(nums, row, record, &result)
	return result
}

func backtracking(nums, row []int, record []uint8, result *[][]int) {
	if len(row) == len(nums) {
		tmp := make([]int, len(row))
		copy(tmp, row)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if record[i] == 1 {
			continue
		}
		row = append(row, nums[i])
		record[i] = 1
		backtracking(nums, row, record, result)
		row = row[:len(row)-1]
		record[i] = 0
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
