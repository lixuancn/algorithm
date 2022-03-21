package main

import "fmt"

//背包问题讲解的练习题 完全背包。完全背包是物品的数量无限制。01背包是只有每样物品只有1个，要么放要么不放
//https://www.programmercarl.com/%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%80%E5%AE%8C%E5%85%A8%E8%83%8C%E5%8C%85.html
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	weight = []int{1, 4, 3}
	value = []int{15, 30, 20}
	fmt.Println(test_CompletePack_1(weight, value, 4))
	fmt.Println(test_CompletePack_2(weight, value, 4))
	fmt.Println(test_CompletePack_3(weight, value, 4))
}

func test_CompletePack_1(weight, value []int, bagCap int) int {
	//dp[i][j]表示物品i放入容量为j的背包中的最大价值
	dp := make([][]int, len(value))
	for i := range dp {
		dp[i] = make([]int, bagCap+1)
	}
	for j := 0; j <= bagCap; j++ {
		dp[0][j] = value[0] * (j / weight[0])
	}
	//先循环i和j都可以的
	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagCap; j++ {
			//要么不放和之前一样（dp[i-1][j]）
			//要么放，那就是在背包容量为j-weight[i]时的最大值，加上i的价值。如包容量4，i重量为1，那么放i的最大值就是包为3时的最大值+i的价值
			//取最大值
			dp[i][j] = dp[i-1][j]
			if j >= weight[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagCap]
}

//对方法1的优化，对dp降维
func test_CompletePack_2(weight, value []int, bagCap int) int {
	//dp[j]表示背包容量为j时所装的最大价值
	//放入i时，dp[j] = dp[j-weight[i]]+value[i]
	//不放i时，dp[j]就还是dp[j]
	//dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
	dp := make([]int, bagCap+1)
	dp[0] = 0
	for i := 0; i < len(weight); i++ {
		//因为完全背包问题可以放多个，所以需要正序遍历
		for j := weight[i]; j <= bagCap; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagCap]
}

//一维dp中，i和j是可以互换位置的
func test_CompletePack_3(weight, value []int, bagCap int) int {
	//dp[j]表示背包容量为j时所装的最大价值
	//放入i时，dp[j] = dp[j-weight[i]]+value[i]
	//不放i时，dp[j]就还是dp[j]
	//dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
	dp := make([]int, bagCap+1)
	dp[0] = 0
	for j := 0; j <= bagCap; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
	return dp[bagCap]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
