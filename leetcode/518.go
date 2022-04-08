package main

import (
	"fmt"
)

//518. 零钱兑换 II

func main() {
	r := change(5, []int{1, 2, 5})
	fmt.Println(r)
}

func change(amount int, coins []int) int {
	//return change_1(amount, coins) //题解
	return change_2(amount, coins) //代码随想录 完全背包问题，一维dp
}

//动态规划 - 背包问题
//核心问题是求不同硬币的组合数。
//dp[i]表示，金额和为i时的组合数，最终结果是求dp[amount]
//当循环到一个coin硬币时，此时的组合数为(i-coin)的组合数，这时把coin接收了之后正好就等于i了嘛。i的范围是从coin到amount
//第二次循环到i时，新组合数是(i-coin)，还得加上之前循环到i时的组合数，即dp[i]=dp[i]+dp[i-coin]
func change_1(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	return dp[amount]
}

func change_2(amount int, coins []int) int {
	//转换为完全背包问题，amount是背包容量，物品价值和重量都是coins[i]
	//dp[j] 表示容量为j的背包的最大组合数
	dp := make([]int, amount+1)
	dp[0] = 1
	//这里i和j不能换位置，因为这里求组合而不是排列，假设coins=1, coins[1]=5
	//i在外时，1会计算一次，5会计算一次。所以只有{1,5}而不会有{5,1}，所以是组合
	//i在内时，{1,5}和{5，1}都会有，所以是排列
	//如果求组合数就是外层for循环遍历物品，内层for遍历背包。
	//如果求排列数就是外层for遍历背包，内层for循环遍历物品。
	//如果把遍历nums（物品）放在外循环，遍历target的作为内循环的话，举一个例子：计算dp[4]的时候，结果集只有 {1,3} 这样的集合，不会有{3,1}这样的集合，因为nums遍历放在外层，3只能出现在1后面！
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func change_practice(amount int, coins []int) int {
	//dp[i]表示背包容量为i时的组合数
	dp := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[amount]
}
