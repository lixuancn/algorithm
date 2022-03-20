package main

import "fmt"

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used := make([]bool, len(nums))
	var backtracking func(k int, bucket int, nums []int, start int, used []bool, target int) bool
	backtracking = func(k int, bucket int, nums []int, start int, used []bool, target int) bool {
		if k == 0 {
			return true
		}
		if bucket == target {
			return backtracking(k-1, 0, nums, 0, used, target)
		}
		for i := start; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if bucket+nums[i] > target {
				continue
			}
			used[i] = true
			bucket += nums[i]
			if backtracking(k, bucket, nums, i+1, used, target) {
				return true
			}
			used[i] = false
			bucket -= nums[i]
		}
		return false
	}
	result := backtracking(k, 0, nums, 0, used, target)
	return result
}

func main() {
	//fmt.Println(canPartitionKSubsets([]int{5, 5, 5, 5}, 2))
	//fmt.Println(canPartitionKSubsets([]int{1, 5, 11, 5}, 2))
	//fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	//fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 3, 4, 4, 4, 5, 6}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1}, 4))
}
