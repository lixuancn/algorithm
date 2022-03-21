package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findTargetSumWays_dp1([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays_dp2([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays_dp3([]int{1, 1, 1, 1, 1}, 3)) //代码随想录
	fmt.Println(findTargetSumWays_recursion([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	return findTargetSumWays_dp1(nums, target)
	return findTargetSumWays_dp2(nums, target)
	return findTargetSumWays_dp3(nums, target)
	return findTargetSumWays_recursion(nums, target)
}

//动态规划优化版
//每次循环只和上一行有关，所以去掉数组第一维。
func findTargetSumWays_dp1(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for i := range nums {
		for j := neg; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[neg]
}

//动态规划
//元素和为sum，加-号的元素和是neg，加+号的元素和就是sum-neg。
//target = (sum - neg) - neg = sum - 2*neg 则 neg = (sum-target)/2
//nums中都是非负整数，所以neg也是非负整数，则sum-target是非负偶数。
//问题转化为：在nums中选取若干元素，使元素和为neg，求方案数。可用动态规划
//定义dp[i][j]，表示nums的前i个元素中选择若干元素，元素和为j的组合数。
//最终所需结果为dp[n][neg]
//边界，0个元素时，和为0，方案数为1。dp[0][0] = 1
//1<=i<=n时，遍历0<=j<=neg计算dp[i][j]：
//	如果j<num，则不能选num，dp[i][j] = dp[i-1][j]
//	如果j>num，如果不选num，dp[i][j] = dp[i-1][j]，如果选，dp[i][j] = dp[i-1][j]+dp[i-1][j-num]
func findTargetSumWays_dp2(nums []int, target int) int {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	dp := make([][]int, n+1)
	dp[0] = make([]int, neg+1)
	dp[0][0] = 1
	for i := range nums {
		dp[i+1] = make([]int, neg+1)
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[n][neg]
}

//01背包问题，left+right = sum      left-right=target ->  left-(sum-left)=target -> left=(target+sum)/2
//转化为挑选背包容量为neg的元素，即和为neg的时候有几种选择法
func findTargetSumWays_dp3(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if target > sum {
		return 0
	}
	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (target+sum)%2 != 0 {
		return 0
	}
	neg := (sum + target) / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := neg; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[neg]
}

//递归
var count int

func findTargetSumWays_recursion(nums []int, target int) int {
	count = 0
	recursion(nums, target, len(nums), 0, nums[0])
	recursion(nums, target, len(nums), 0, -nums[0])
	return count
}

func recursion(nums []int, target, n, index, sum int) {
	if index == n-1 {
		if sum == target {
			count++
		}
		return
	}
	i := index + 1
	recursion(nums, target, n, i, sum+nums[i])
	recursion(nums, target, n, i, sum-nums[i])
}
