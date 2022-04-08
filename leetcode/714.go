package main

import "fmt"

//714. 买卖股票的最佳时机含手续费

//dp[i][j]表示第i天的j状态下的最大收益，j是0=现金，1=持股
func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[len(prices)-1][0]
}

//贪心 很难理解
//负数不要，正数就是收益。思路和122. 买卖股票的最佳时机 II 是不一样的 不 一 样。
//因为可能涉及到手续费，每天算利润的话，就相当于每天收手续费，而现实是这天可能并不真的卖
func maxProfit_1(prices []int, fee int) int {
	ret, buyPoint := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		//情况二：昨天是收获日，真正的卖出了。今天重新计算买点
		if prices[i] < buyPoint {
			buyPoint = prices[i]
		}
		//情况三：保持不动，今天卖能赚，但加上手续费就亏了
		if prices[i] >= buyPoint && prices[i]-buyPoint < fee {
			continue
		}

		//情况一：收获的这一天并不是最后一天，也就是并不是真的卖出
		//计算利润, 可能没有真正卖
		if prices[i] > buyPoint+fee {
			ret += prices[i] - buyPoint - fee
			//更新买点（如果还在收获利润的区间里，表示并不是真正的卖出，而计算利润每次都要减去手续费，所以要让minBuy = prices[i] - fee;，这样在明天收获利润的时候，才不会多减一次手续费！）
			buyPoint = prices[i] - fee //手续费第一次已经算过了
		}
	}
	return ret
}

func maxProfit_dp_practice(prices []int, fee int) int {
	//dp[i][0]现金，dp[i][1]股票
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, -prices[0]-fee
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		//这里手续费要放在买入的地方，这样才可以和初始化dp[0]的时候对应
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[len(prices)-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit_1([]int{1, 3, 2, 8, 4, 9}, 2))
}
