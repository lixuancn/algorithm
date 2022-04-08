package main

import (
	"fmt"
)

//322. 零钱兑换

//零钱兑换 背包问题，物品数量不限，所以时完全背包问题，amount是背包容量。
//因为是组合，所以外层循环物品，内层循环背包容量。不能替换位置。
//dp[j]表示背包容量为j时，凑足j所需要的最少硬币个数

//如果求组合数就是外层for循环遍历物品，内层for遍历背包。
//如果求排列数就是外层for遍历背包，内层for循环遍历物品。
//如果把遍历nums（物品）放在外循环，遍历target的作为内循环的话，举一个例子：计算dp[4]的时候，结果集只有 {1,3} 这样的集合，不会有{3,1}这样的集合，因为nums遍历放在外层，3只能出现在1后面！

//本题求最小数，所以组合/排列都影响结果，所以循环顺序都可以
//完全背包，所以内循环正序
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1<<31 - 1
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			//放i时，dp[j-coins[i]] + 1
			//不放时，dp[j]
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == 1<<31-1 {
		return -1
	}
	return dp[amount]
}

//完全背包问题
func coinChange_practice(coins []int, amount int) int {
	//dp[i]表示背包容量为时，填满需要的最小个数
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1 << 31
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount; i++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == 1<<31 {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
