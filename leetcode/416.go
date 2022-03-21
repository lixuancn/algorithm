package main

import (
	"fmt"
	"sort"
)

//动态规划 01背包问题 重量是nums，价值时nums数组的物品放入背包，价值恰好时数组和的一半
//dp[j]表示容量为j时最大的子集和
func canPartition(nums []int) bool {
	bagCap := 0
	for _, num := range nums {
		bagCap += num
	}
	if bagCap&1 == 1 {
		return false
	}
	bagCap /= 2
	dp := make([]int, bagCap+1)
	for i := 0; i < len(nums); i++ {
		for j := bagCap; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[bagCap] == bagCap
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//递归 回溯 会超时
func canPartition_recursion(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var backtracking func(nums []int, k, bucket, target, start int, used []bool) bool
	backtracking = func(nums []int, k, bucket, target, start int, used []bool) bool {
		if k == 0 {
			return true
		}
		if bucket > target {
			return false
		}
		if bucket == target {
			return backtracking(nums, k-1, 0, target, 0, used)
		}
		for i := start; i < len(nums); i++ {
			if bucket+nums[i] > target {
				break
			}
			if used[i] {
				continue
			}
			bucket += nums[i]
			used[i] = true
			if backtracking(nums, k, bucket, target, i+1, used) {
				return true
			}
			bucket -= nums[i]
			used[i] = false
		}
		return false
	}
	return backtracking(nums, 2, 0, sum, 0, used)
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
