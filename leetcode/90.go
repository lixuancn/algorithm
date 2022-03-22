package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	row := make([]int, 0)
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	backtracking(nums, 0, row, &result, used)
	return result
}

func backtracking(nums []int, start int, row []int, result *[][]int, used []bool) {
	tmp := make([]int, len(row))
	copy(tmp, row)
	*result = append(*result, tmp)
	for i := start; i < len(nums); i++ {
		//去重
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		// 而我们要对同一树层使用过的元素进行跳过
		if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
			continue
		}
		//去重也可以这么写
		//if (i > start && nums[i] == nums[i - 1] ) {
		//	continue;
		//}
		row = append(row, nums[i])
		used[i] = true
		backtracking(nums, i+1, row, result, used)
		row = row[:len(row)-1]
		used[i] = false
	}
}
