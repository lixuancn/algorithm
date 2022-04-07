package main

import "fmt"

//70. 爬楼梯

func climbStairs(n int) int {
	//dp[i] = dp[i-1] + dp[i-2]
	if n <= 2 {
		return n
	}
	prev1, prev2, result := 2, 1, 0 //i-1   n-2   result
	for i := 2; i < n; i++ {
		result = prev1 + prev2
		prev2 = prev1
		prev1 = result
	}
	return result
}

func climbStairs_practice(n int) int {
	//dp[i]表示爬到第i阶楼梯一共有多少种方法
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func main() {
	fmt.Println(climbStairs_practice(2))
	fmt.Println(climbStairs_practice(3))
	fmt.Println(climbStairs_practice(4))
}
