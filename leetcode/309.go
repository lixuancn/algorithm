package main

import "fmt"

//309. 最佳买卖股票时机含冷冻期
//dp[i][j] 表示第i天 j状态下的最大现金
//状态0：买入或持股中
//状态1：之前卖出的，且不在冷冻期，保持空仓状态
//状态2：今天卖出
//状态3：昨天卖出所以今天是冷冻期
//dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]) 保持持股、昨天是冷冻期今天买入，之前空仓今天买入
//dp[i][1] = max(dp[i-1][1], dp[i-1][3]) 保持空仓和之前一样、昨天是冷冻期
//dp[i][2] = dp[i-1][0]+prices[i] //昨天的持股状态加上今天的售价
//dp[i][3] = dp[i-1][2] //昨天卖出的
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][4]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	dp[0][3] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(max(dp[i-1][0], dp[i-1][1]-prices[i]), dp[i-1][3]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(max(dp[len(prices)-1][1], dp[len(prices)-1][2]), dp[len(prices)-1][3])
}

//空间优化
func maxProfit_1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	cur0, prev0 := -prices[0], -prices[0]
	cur1, prev1 := 0, 0
	cur2, prev2 := 0, 0
	cur3, prev3 := 0, 0
	for i := 1; i < len(prices); i++ {
		cur0 = max(max(prev0, prev1-prices[i]), prev3-prices[i])
		cur1 = max(prev1, prev3)
		cur2 = prev0 + prices[i]
		cur3 = prev2
		prev0 = cur0
		prev1 = cur1
		prev2 = cur2
		prev3 = cur3
	}
	return max(max(prev1, prev2), prev3)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit_1([]int{1, 2, 3, 0, 2}))
}
