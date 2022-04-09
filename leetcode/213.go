package main

import "fmt"

//213. 打家劫舍 II

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	//dp[i]表示前i个房屋可以偷到的最大值
	//偷的话，dp[i] = dp[i-2]+nums[i]，不偷的话dp[i]=dp[i-1]
	n := len(nums)
	//不考虑尾
	dp1 := make([]int, n-1)
	dp1[0], dp1[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n-1; i++ {
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
	}
	//不考虑头
	dp2 := make([]int, n-1)
	dp2[0], dp2[1] = nums[1], max(nums[1], nums[2])
	for i := 2; i < n-1; i++ {
		dp2[i] = max(dp2[i-2]+nums[i+1], dp2[i-1])
	}
	return max(dp1[n-2], dp2[n-2])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(rob([]int{1, 1}))
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
