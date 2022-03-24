package main

import "fmt"

//188. 买卖股票的最佳时机 IV

//dp[i][j]表示，第i天的j状态下所持有的的最大现金
//思路和 123. 买卖股票的最佳时机 III 类似
//j有五种状态：0无操作，1第一次买入或持股中，2第一次卖出或保持空仓，3第二次买入或持股中，4第二次卖出或保持空仓
//dp[i][0] = dp[i-1][0]
//dp[i][1] 和昨天一样保持或用昨天的现金买入 = max(dp[i-1][1], dp[i-1][0]-prices[i])
//dp[i][2] 和昨天一样保持或把昨天持股卖出 = max(dp[i-1][2], dp[i-1][1]+prices[i])
//dp[i][3] 和昨天一样保持或用昨天空仓时的现金买入 = max(dp[i-1][3], dp[i-1][2]-prices[i])
//dp[i][4] 和昨天一样保持或用昨天的持股卖出 = max(dp[i-1][4], dp[i-1][3]+prices[i])
//..........第N次买卖..........
//可以发现，j为0都是和前一个一样
//j为奇数都是买
//j为偶数都是卖
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 1+2*k) //一个j=0，k次买和k次卖
	}
	//初始化，i=0的情况
	dp[0][0] = 0
	for j := 1; j < 2*k+1; j++ {
		if j&1 == 1 {
			dp[0][j] = -prices[0]
		} else {
			dp[0][j] = 0
		}
	}
	//开始循环
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k+1; j++ {
			if j&1 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit_1(2, []int{3, 2, 6, 5, 0, 3}))
}
