package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			//i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			//因为同一个i有不同分拆组合，所以要取最大的dp[i]
			dp[i] = max(dp[i], max((i-j)*j, j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(3))
	fmt.Println(integerBreak(10))
}
