package main

import "fmt"

func superEggDrop(k int, n int) int {
	return dp_force(k, n) //动态规划-暴力法
	//return dp_mid(k, n) //动态规划-二分法 - 未完成
}

//dp[i][j]表示一共有i层时，一共有j个鸡蛋时，确定F的最小次数。i表示的楼层总数，不是第几层，比如[7,8,9]时n为3
//如果在x层碎了，就需要dp[i][j]=dp[x-1][j-1]次，没碎的话，就需要dp[i][j]=dp[i-x][j]次，取最大值
//对于不同的x，则取最小值
func dp_force(k, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	//初始化
	for i := 0; i < n+1; i++ {
		for j := 0; j < k+1; j++ {
			//结果需要求最小值，所以初始化成最大值
			dp[i][j] = 1 << 31
			//第0层时无法做实验，所以dp为0
			if i == 0 {
				dp[0][j] = 0
			}
			//第1层时，0个鸡蛋需要扔0次，1个以上的鸡蛋需要扔1从
			if i == 1 {
				if j == 0 {
					dp[1][0] = 0
				} else {
					dp[1][j] = 1
				}
			}
			//鸡蛋数为0的时候，不管多少层都无法确认
			if j == 0 {
				dp[i][0] = 0
			}
			//鸡蛋数为1的时候，需要从1层挨着往上走，才能找到F，所以需要n层
			if j == 1 {
				dp[i][1] = i
			}
		}
	}
	//动态规划
	for i := 2; i < n+1; i++ {
		for j := 2; j < k+1; j++ {
			//i层高的楼，开始尝试
			for k := 1; k <= i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k-1][j-1], dp[i-k][j])+1)
			}
		}
	}
	return dp[n][k]
}

//未完成

func dp_mid(k, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	//初始化
	for i := 0; i < n+1; i++ {
		for j := 0; j < k+1; j++ {
			//结果需要求最小值，所以初始化成最大值
			dp[i][j] = 1 << 31
			//第0层时无法做实验，所以dp为0
			if i == 0 {
				dp[0][j] = 0
			}
			//第1层时，0个鸡蛋需要扔0次，1个以上的鸡蛋需要扔1从
			if i == 1 {
				if j == 0 {
					dp[1][0] = 0
				} else {
					dp[1][j] = 1
				}
			}
			//鸡蛋数为0的时候，不管多少层都无法确认
			if j == 0 {
				dp[i][0] = 0
			}
			//鸡蛋数为1的时候，需要从1层挨着往上走，才能找到F，所以需要n层
			if j == 1 {
				dp[i][1] = i
			}
		}
	}
	//动态规划
	for i := 2; i < n+1; i++ {
		for j := 2; j < k+1; j++ {
			left, right := 1, i
			for left < right {
				mid := (left + right) >> 1
				if dp[mid-1][j-1] > dp[i-mid][j] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			dp[i][j] = min(dp[i][j], max(dp[left-1][j-1], dp[i-left][j])+1)
		}
	}
	return dp[n][k]
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
	//fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))
}
