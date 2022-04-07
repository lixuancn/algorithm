package main

import (
	"fmt"
)

//746. 使用最小花费爬楼梯

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	//从第i个台阶出发，所需要的钱
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minCostClimbingStairs_practice(cost []int) int {
	//dp[i]表示上第i阶楼梯的最小费用
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[len(cost)-2], dp[len(cost)-1])
}

func main() {
	fmt.Println(minCostClimbingStairs_practice([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs_practice([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
