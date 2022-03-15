package main

import "fmt"

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

func main() {
	fmt.Println(climbStairs(4))
}
