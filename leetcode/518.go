package main

import (
	"fmt"
)

func main() {
	r := change(5, []int{1,2,5})
	fmt.Println(r)
}

//动态规划 - 背包问题
//核心问题是求不同硬币的组合数。
//dp[i]表示，金额和为i时的组合数，最终结果是求dp[amount]
//当循环到一个coin硬币时，此时的组合数为(i-coin)的组合数，这时把coin接收了之后正好就等于i了嘛。i的范围是从coin到amount
//第二次循环到i时，新组合数是(i-coin)，还得加上之前循环到i时的组合数，即dp[i]=dp[i]+dp[i-coin]
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins{
		for i:=coin; i<=amount; i++{
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	fmt.Println(dp)
	return dp[amount]
}
