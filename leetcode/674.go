package main

import "fmt"

//674. 最长连续递增序列

func findLengthOfLCIS(nums []int) int {
	//return findLengthOfLCIS_recursion(nums) //回溯
	return findLengthOfLCIS_dp_practice(nums) //动态规划
	return findLengthOfLCIS_tanxin(nums)      //贪心
}

var max = 0

func findLengthOfLCIS_recursion(nums []int) int {
	max = 0
	row := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		backtracking(nums, i, row)
	}
	return max
}

func backtracking(nums []int, start int, row []int) bool {
	if max < len(row) {
		max = len(row)
	}
	for i := start; i < len(nums); i++ {
		if len(row) > 0 && row[len(row)-1] >= nums[i] {
			return false
		}
		row = append(row, nums[i])
		result := backtracking(nums, i+1, row)
		row = row[:len(row)-1]
		if !result {
			return result
		}
	}
	return true
}

//dp[i]表示放入i时的最长连续递增子序列的长度
func findLengthOfLCIS_dp(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func findLengthOfLCIS_dp_practice(nums []int) int {
	//dp[i]表示放入nums[i]时的连续最长子序列。
	//nums[i] > nums[i-1]时，dp[i] = dp[i-1]+1
	result := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func findLengthOfLCIS_tanxin(nums []int) int {
	max := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		if max < count {
			max = count
		}
	}
	return max
}

func findLengthOfLCIS_tanxin_practice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			length++
		} else {
			length = 1
		}
		if length > result {
			result = length
		}
	}
	return result
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
}
