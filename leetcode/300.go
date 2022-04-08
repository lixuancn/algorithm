package main

import (
	"fmt"
)

//300. 最长递增子序列

func lengthOfLIS(nums []int) int {
	//return lengthOfLIS_recursion(nums)
	return lengthOfLIS_dp(nums)
}

//递归
func lengthOfLIS_recursion(nums []int) int {
	max := 0
	row := make([]int, 0)
	backtracking(nums, row, 0, &max)
	return max
}

func backtracking(nums, row []int, start int, max *int) {
	if len(row) > *max {
		*max = len(row)
	}
	for i := start; i < len(nums); i++ {
		//因为要递增
		if len(row) > 0 && nums[i] <= row[len(row)-1] {
			continue
		}
		row = append(row, nums[i])
		backtracking(nums, row, i+1, max)
		row = row[:len(row)-1]
	}
}

//动态规划
func lengthOfLIS_dp(nums []int) int {
	//dp[i]表示放入nums[i]时的子串长度（以nums[i]结尾）
	//0<j<i，那么nums[i] > nums[j]时，可以放入nums[i]，则dp[i] = dp[j]+1
	//nums[i] <= nums[j] 无法放入，跳过
	max := 0
	dp := make([]int, len(nums))
	//每个元素自己就是一个子序列
	for k, _ := range dp {
		dp[k] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				//dp[i] = max(dp[i], dp[j]+1)
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func lengthOfLIS_dp_practice(nums []int) int {
	//dp[i]是[0,i]子数组中的最长递增子序列。有0<j<i
	//如果nums[i] > nums[j]，则可以放入i，那么此时子序列长度是dp[j]+1
	//如果nums[i] < nums[j]，则无法放入，跳过
	//dp[i] = max(dp[i], dp[j]+1)
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; i++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	//fmt.Println(lengthOfLIS([]int{7, 7, 7, 7}))
}
