package main

import (
	"fmt"
)

func main() {
	//fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
	fmt.Println(lastStoneWeightII([]int{2, 4, 1, 1}))
}

func lastStoneWeightII(stones []int) int {
	return lastStoneWeightII_dp1(stones)
	//return lastStoneWeightII_dp2(stones)
	return lastStoneWeightII_dp3(stones) //动态规划 一维 代码随想录
}

//动态规划 - 背包问题 和494类似
//题目翻译(a-b)+(b-c)最小。
//sum表示stones所有元素之和，neg表示所有为负的元素和，为正的元素和就是sum-neg
//粉碎后剩余的石头为(sum-neg)-neg = sum - 2*neg。 剩余石头最小时neg要无限接近于sum/2
//背包问题：背包容量为sum/2，物品重量和价值均为stones[i]的0-1背包问题。
//dp[i][j]表示前i个元素，剩余重量是否为j。(前i个石头能否凑出重量j)
//边界：i=0时，j=0为真，其余j均为假
//	j<stones[i]，则不能选，dp[i][j] = dp[i-1][j]
//  j>=stones[i]，若不选，dp[i][j] = dp[i-1][j]；若选，dp[i][j] = dp[i-1][j-stones[i]]。看这种方法谁为真。
//遍历dp[n]，最大的j就是答案
func lastStoneWeightII_dp1(stones []int) int {
	sum, n := 0, len(stones)
	for _, stone := range stones {
		sum += stone
	}
	diff := sum / 2
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, diff+1)
	}
	dp[0][0] = true
	for i, stone := range stones {
		for j := 0; j <= diff; j++ {
			fmt.Println(i, j)
			dp[i+1][j] = dp[i][j]
			if j >= stone {
				dp[i+1][j] = dp[i+1][j] || dp[i][j-stone]
			}
			fmt.Println(dp)
		}
	}
	for j := diff; j >= 0; j-- {
		if dp[n][j] {
			return sum - 2*j
		}
	}
	return 0
}

//动态规划优化
//由于dp[i][j]只和上一行相关，则可以把第一维去掉
func lastStoneWeightII_dp2(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	diff := sum / 2
	dp := make([]bool, diff+1)
	dp[0] = true
	for _, stone := range stones {
		for j := diff; j >= stone; j-- {
			dp[j] = dp[j] || dp[j-stone]
		}
	}
	for j := diff; j >= 0; j-- {
		if dp[j] {
			return sum - 2*j
		}
	}
	return 0
}

//01背包问题，重量和价值都是stones[i]，求哪个组合更接近sum/2，这样碰撞就更小
//dp[j]表示容量为j的背包所能装入的最大价值
//如果不装，dp[j] = dp[j]
//如果装，dp[j] = dp[j-stones[i]]+stones[i]
//记得j必须倒序
func lastStoneWeightII_dp3(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			//和max(dp[j], dp[j-stones[i]]+stones[i])是一样的
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	//剩余的一推石头价值 - 我们跳出来的石头价值
	return (sum - dp[target]) - dp[target]
}
