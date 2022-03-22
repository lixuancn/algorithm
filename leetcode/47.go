package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	row := make([]int, 0)
	result := make([][]int, 0)
	record := make([]bool, len(nums))
	sort.Ints(nums)
	backtracking(nums, row, record, &result)
	return result
}

func backtracking(nums []int, row []int, record []bool, result *[][]int) {
	if len(row) == len(nums) {
		tmp := make([]int, len(row))
		copy(tmp, row)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		//这是排重，抽象成树的话，就是同一层的元素不能相同
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		// 如果同一树层nums[i - 1]使用过则直接跳过
		if i > 0 && nums[i-1] == nums[i] && record[i-1] == false {
			continue
		}
		if record[i] {
			continue
		}
		row = append(row, nums[i])
		record[i] = true
		backtracking(nums, row, record, result)
		row = row[:len(row)-1]
		record[i] = false
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
