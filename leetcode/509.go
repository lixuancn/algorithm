package main

import "fmt"

//动态规划 斐波那契数列
func fib_dp(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//动态规划 斐波那契数列
func fib(n int) int {
	if n <= 1 {
		return n
	}
	prev1 := 1 //n-1
	prev2 := 0 //n-2
	result := prev1 + prev2
	for i := 2; i <= n; i++ {
		result = prev1 + prev2
		prev2 = prev1
		prev1 = result
	}
	return result
}

func main() {
	fmt.Println(fib(4))
}
