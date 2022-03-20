package main

import "fmt"

//背包问题讲解的练习题
//https://www.programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.html
//我一直在想可能需要有序，但实践发现，不需要。
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	weight = []int{1, 4, 3}
	value = []int{15, 30, 20}
	fmt.Println(test_2_wei_bag_problem1(weight, value, 4))
}

func test_2_wei_bag_problem1(weight, value []int, bagCap int) int {
	//dp[i][j]表示物品i放入容量为j的背包中的最大价值
	dp := make([][]int, len(value))
	for i := range dp {
		//因为有bagCap=0
		dp[i] = make([]int, bagCap+1)
	}
	for j := 0; j <= bagCap; j++ {
		dp[0][j] = 0
		if weight[0] <= j {
			dp[0][j] = value[0]
		}
	}
	//递推公式
	for i := 1; i < len(value); i++ {
		for j := 0; j <= bagCap; j++ {
			//要么不放和之前一样（dp[i-1][j]）
			//要么放，那就是在背包容量为j-weight[i]时的最大值，加上i的价值。如包容量4，i重量为1，那么放i的最大值就是包为3时的最大值+i的价值
			//取最大值
			dp[i][j] = dp[i-1][j]
			if weight[i] <= j { //包的容量够用
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(weight)-1][bagCap]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
