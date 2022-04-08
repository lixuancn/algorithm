package main

import "fmt"

//123. 买卖股票的最佳时机 III

//dp[i][j]表示，第i天的j状态下所持有的的最大现金
//j有五种状态：0无操作，1第一次买入或持股中，2第一次卖出或保持空仓，3第二次买入或持股中，4第二次卖出或保持空仓
//dp[i][0] = dp[i-1][0]
//dp[i][1] 和昨天一样保持或用昨天的现金买入 = max(dp[i-1][1], dp[i-1][0]-prices[i])
//dp[i][2] 和昨天一样保持或把昨天持股卖出 = max(dp[i-1][2], dp[i-1][1]+prices[i])
//dp[i][3] 和昨天一样保持或用昨天空仓时的现金买入 = max(dp[i-1][3], dp[i-1][2]-prices[i])
//dp[i][4] 和昨天一样保持或用昨天的持股卖出 = max(dp[i-1][4], dp[i-1][3]+prices[i])
func maxProfit(prices []int) int {
	dp := make([][5]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}

//优化空间版
func maxProfit_1(prices []int) int {
	dp := [5]int{}
	dp[0] = 0
	dp[1] = -prices[0]
	dp[2] = 0
	dp[3] = -prices[0]
	dp[4] = 0
	for i := 1; i < len(prices); i++ {
		dp[0] = dp[0]
		dp[1] = max(dp[1], dp[0]-prices[i])
		dp[2] = max(dp[2], dp[1]+prices[i])
		dp[3] = max(dp[3], dp[2]-prices[i])
		dp[4] = max(dp[4], dp[3]+prices[i])
	}
	return dp[4]
}

func maxProfit_practice(prices []int) int {
	//dp[i][0]初始化持有现金 dp[i][1]第一次持股中 dp[i][2]第一次卖出持现金 dp[i][3]第二次持股 dp[i][4]第二次卖出后持现金
	dp := make([][5]int, len(prices))
	dp[0][0], dp[0][1], dp[0][2], dp[0][3], dp[0][4] = 0, -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit_practice([]int{5, 4, 3, 2, 1}))
	//fmt.Println(maxProfit_practice([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
