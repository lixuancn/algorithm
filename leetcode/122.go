package main

import "fmt"

//贪心
//局部最优：收集每天的正利润，全局最优：求得最大利润。
//其实我们需要收集每天的正利润就可以，收集正利润的区间，就是股票买卖的区间，而我们只需要关注最终利润，不需要记录区间。
//所以第一天买第三天卖的过程，可以转换为第一天买第二天卖，第二天买第三天卖，然后2次利润相加
func maxProfit(prices []int) int {
	ret := 0
	for i := 1; i < len(prices); i++ {
		todayProfit := prices[i] - prices[i-1]
		if todayProfit > 0 {
			ret += todayProfit
		}
	}
	return ret
}

//动态规划
func maxProfit_dp(prices []int) int {
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0] //股票
	dp[0][1] = 0          //现金
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit_dp([]int{7, 1, 5, 3, 6, 4}))
}
