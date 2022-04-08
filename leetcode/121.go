package main

import "fmt"

//121. 买卖股票的最佳时机

func maxProfit(prices []int) int {
	//return maxProfit_force(prices) //暴力
	return maxProfit_tanxin(prices) //贪心
	//return maxProfit_dp_1(prices) //动态规划
	return maxProfit_dp_2(prices) //动态规划
}

func maxProfit_force(prices []int) int {
	ret := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > ret {
				ret = prices[j] - prices[i]
			}
		}
	}
	return ret
}

//贪心，因为只能买卖一次，所以想到贪心
func maxProfit_tanxin(prices []int) int {
	buyPoint := prices[0]
	ret := 0
	for i := 1; i < len(prices); i++ {
		buyPoint = min(buyPoint, prices[i])
		ret = max(ret, prices[i]-buyPoint)
	}
	return ret
}

func maxProfit_dp_1(prices []int) int {
	//dp[i][0] 股票
	//dp[i][1] 现金
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		//因为只能买一次，所以比较的后半部分直接就是0-prices[i]，而不是dp[i-1][1]-prices[i]
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func maxProfit_dp_2(prices []int) int {
	//dp[i]表示最大收益
	//dp[i] = max(dp[i-1], prices[i]-minBuy)
	//我觉得思路更像贪心
	dp := make([]int, len(prices))
	minBuy := prices[0]
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		minBuy = min(minBuy, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minBuy)
	}
	return dp[len(prices)-1]
}

func maxProfit_dp_practice(prices []int) int {
	//dp[i][0] 持股 dp[i][1] 持现金
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func maxProfit_tanxin_practice(prices []int) int {
	ans, minBuy := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		minBuy = min(minBuy, prices[i])
		ans = max(ans, prices[i]-minBuy)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
