package main

import (
	"fmt"
)

//暴力
func twoSum_force(nums []int, target int) []int {
	//record := make(map[int]struct{})
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//map
func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for k, v := range nums {
		if index, ok := record[target-v]; ok {
			return []int{index, k}
		}
		record[v] = k
	}
	return nil
}

func twoSum_practice(nums []int, target int) []int {
	haved := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := haved[target-nums[i]]; ok {
			return []int{i, index}
		}
		haved[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
