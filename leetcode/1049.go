package main

import (
	"fmt"
)

func main() {
	r := lastStoneWeightII([]int{31,26,33,21,40})
	fmt.Println(r)
}

//动态规划优化
//由于dp[i][j]只和上一行相关，则可以把第一维去掉
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones{
		sum += stone
	}
	diff := sum / 2
	dp := make([]bool, diff+1)
	dp[0] = true
	for _, stone := range stones{
		for j:=diff; j>=stone; j--{
			dp[j] = dp[j] || dp[j-stone]
		}
	}
	for j:=diff; j>=0; j--{
		if dp[j]{
			return sum - 2 * j
		}
	}
	return 0
}

//动态规划 - 背包问题 和494类似
//题目翻译(a-b)+(b-c)最小。
//sum表示stones所有元素之和，neg表示所有为负的元素和，为正的元素和就是sum-neg
//粉碎后剩余的石头为(sum-neg)-neg = sum - 2*neg。 剩余石头最小时neg要无限接近于sum/2
//背包问题：背包容量为sum/2，物品重量和价值均为stones[i]的0-1背包问题。
//dp[i][j]表示前i个元素，剩余重量是否为j。
//边界：i=0时，j=0为真，其余j均为假
//	j<stones[i]，则不能选，dp[i][j] = dp[i-1][j]
//  j>=stones[i]，若不选，dp[i][j] = dp[i-1][j]；若选，dp[i][j] = dp[i-1][j-stones[i]]。看这种方法谁为真。
//遍历dp[n]，最大的j就是答案
func lastStoneWeightII_dp(stones []int) int {
	sum, n := 0, len(stones)
	for _, stone := range stones{
		sum += stone
	}
	diff := sum / 2
	dp := make([][]bool, n+1)
	for i := range dp{
		dp[i] = make([]bool, diff+1)
	}
	dp[0][0] = true
	for i, stone := range stones{
		for j:=0; j<=diff; j++{
			dp[i+1][j] = dp[i][j]
			if j >= stone{
				dp[i+1][j] = dp[i+1][j] || dp[i][j-stone]
			}
		}
	}
	for j:=diff; j>=0; j--{
		if dp[n][j]{
			return sum - 2 * j
		}
	}
	return 0
}
