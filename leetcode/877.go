package main

import "fmt"

func main() {
	r := stoneGame([]int{5,3,4,5})
	fmt.Println(r)
}

//数学法
func stoneGame(piles []int) bool {
	return true
}

//动态规划，优化空间复杂度
func stoneGame_dp_2(piles []int) bool {
	n := len(piles)
	dp := make([]int, n)
	for i:=0; i<n; i++{
		dp[i] = piles[i]
	}
	for i:=n-2; i>=0; i--{
		for j:=i+1; j<n; j++{
			dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[n-1] > 0
}

//动态规划，dp[i][j]表示从i到j的子石头堆中，俩人总数差值
func stoneGame_dp(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i:=0; i<n; i++{
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for i:=n-2; i>=0; i--{
		for j:=i+1; j<n; j++{
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}

func max(i, j int)int{
	if i > j {
		return i
	}
	return j
}

//递归
func stoneGame_recursion(piles []int) bool {
	return recursion(piles, len(piles), 0)
}

func recursion(piles []int, n, diff int)bool{
	if n == 0{
		if diff > 0{
			return true
		}
		return false
	}

	r := recursion(piles[2:n], n-2, diff+piles[0]-piles[1])
	if r{
		return true
	}
	r = recursion(piles[:n-2], n-2, diff+piles[n-1]-piles[n-2])
	if r{
		return true
	}
	r = recursion(piles[1:n-1], n-2, diff+piles[0]-piles[n-1])
	if r{
		return true
	}
	r = recursion(piles[1:n-1], n-2, diff+piles[n-1]-piles[0])
	if r{
		return true
	}
	return r
}
