package main

import "fmt"

//198. 打家劫舍

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	//dp[i]表示前i个房屋中能偷取到的最大金额，[0]是偷，[1]是不偷
	dp := make([][2]int, len(nums))
	dp[0][0], dp[0][1] = nums[0], 0
	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
	//合并成一维数组
	//dp[i]表示前i个房间偷取到的最大值。偷的话dp[i] = dp[i-2]+nums[i]。不偷的话，dp[i] = dp[i-1]
	dp1 := make([]int, len(nums))
	dp1[0], dp1[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}
	return dp1[len(nums)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))    //4
	fmt.Println(rob([]int{2, 7, 9, 3, 1})) //12
}
