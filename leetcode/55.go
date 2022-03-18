package main

import "fmt"

func canJump(nums []int) bool {
	//return canJump_dp(nums) //动态规划
	//return canJump_recursion_1(nums) //递归 回溯 倒着走
	return canJump_tanxin_1(nums) //贪心 正走
	return canJump_tanxin_2(nums) //贪心 倒走
}

func canJump_dp(nums []int) bool {
	//dp[i]表示跳到下标i时还剩余的步数
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if dp[0] == 0 {
		if len(nums) == 1 {
			return true
		}
		return false
	}
	for i := 1; i < len(nums)-1; i++ {
		dp[i] = max(dp[i-1]-1, nums[i])
		if dp[i] == 0 {
			return false
		}
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//倒着走
func canJump_recursion_1(nums []int) bool {
	var backtracking func([]int, int) bool
	backtracking = func(nums []int, start int) bool {
		if start == 0 {
			return true
		}
		for i := start - 1; i >= 0; i-- {
			if start-i <= nums[i] {
				return backtracking(nums, start-1)
			}
		}
		return false
	}
	return backtracking(nums, len(nums)-1)
}

//正着走
func canJump_tanxin_1(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

//反着走
func canJump_tanxin_2(nums []int) bool {
	end := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= end {
			end = i
		}
	}
	return end == 0
}

func main() {
	fmt.Println(canJump([]int{2, 5, 0, 0}))
	//fmt.Println(canJump([]int{2, 0}))
	//fmt.Println(canJump([]int{0, 1}))
	//fmt.Println(canJump([]int{0}))
	//fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	//fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
