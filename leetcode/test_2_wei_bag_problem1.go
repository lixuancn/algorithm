package main

import "fmt"

//背包问题讲解的练习题 01背包
//https://www.programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.html
//我一直在想可能需要有序，但实践发现，不需要。
func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	weight = []int{1, 4, 3}
	value = []int{15, 30, 20}
	fmt.Println(test_2_wei_bag_problem1(weight, value, 4))
	fmt.Println(test_2_wei_bag_problem1_2(weight, value, 4))
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
	//先循环i和j都可以
	for i := 1; i < len(value); i++ {
		for j := 0; j <= bagCap; j++ {
			//要么不放和之前一样（dp[i-1][j]）
			//要么放，那就是在背包容量为j-weight[i]时的最大值，加上i的价值。如包容量4，i重量为1，那么放i的最大值就是包为3时的最大值+i的价值
			//取最大值

			//包容量步够时选择不放
			dp[i][j] = dp[i-1][j]
			//包的容量够用
			if weight[i] <= j {
				//两种写法都可
				//dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
				dp[i][j] = max(dp[i][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagCap]
}

//对方法1的优化，对dp降维
func test_2_wei_bag_problem1_2(weight, value []int, bagCap int) int {
	//dp[j]表示背包容量为j时所装的最大价值
	//放入i时，dp[j] = dp[j-weight[i]]+value[i]
	//不放i时，dp[j]就还是dp[j]
	//dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
	dp := make([]int, bagCap+1)
	dp[0] = 0
	for i := 0; i < len(weight); i++ {
		//这里j必须倒序，因为j正序时，会把物品放入多次，比如dp[1]=dp[1-1]+value[1]=15, dp[2]=dp[2-1]+value[1]=30
		//这个i和j也不能换，i必须在外层，j在内层，因为j必须倒序，如果更换了的话，背包只会放1个物品。
		for j := bagCap; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
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
