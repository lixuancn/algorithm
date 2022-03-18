package main

import "fmt"

func jump(nums []int) int {
	return jump_1(nums)  // 贪心 反向
	return jump_1(nums)  // 贪心 正向
	return jump_dp(nums) // 动态规划
}

//反向
func jump_1(nums []int) int {
	start := len(nums) - 1
	step := 0
	for start > 0 {
		for i := 0; i < start; i++ {
			if nums[i]+i >= start {
				step++
				start = i
				break
			}
		}
	}
	return step
}

//正向
//在具体的实现中，我们维护当前能够到达的最大下标位置，记为边界。我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
//思想就一句话：每次在上次能跳到的范围（end）内选择一个能跳的最远的位置（也就是能跳到max_far位置的点）作为下次的起跳点 ！
func jump_2(nums []int) int {
	end := 0
	max := 0
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		if max < nums[i]+i {
			max = nums[i] + i
		}
		if end == i {
			end = max
			result++
		}
	}
	return result
}

//dp[i]表示到i点时的最小步
func jump_dp(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 10001
	}
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				//dp[i] = min(dp[i], dp[j]+1)
				if dp[i] > dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[len(nums)-1]
}

func main() {
	fmt.Println(jump_dp([]int{2, 3, 1, 1, 4}))
}
