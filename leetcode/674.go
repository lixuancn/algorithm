package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	//return findLengthOfLCIS_recursion(nums) //回溯 未成功
	return findLengthOfLCIS_dp(nums)     //动态规划
	return findLengthOfLCIS_tanxin(nums) //贪心
}

var max = 0

func findLengthOfLCIS_recursion(nums []int) int {
	max = 0
	row := make([]int, 0)
	backtracking(nums, 0, row)
	return max
}

func backtracking(nums []int, start int, row []int) bool {
	fmt.Println(row)
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

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5}))
	//fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2}))
	//fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
}
