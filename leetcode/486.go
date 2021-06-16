package main

import "fmt"

func main() {
	r := PredictTheWinner([]int{1,5,2})
	fmt.Println(r)
}

//动态规划
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	for i:=0; i<n; i++{
		dp[i] = nums[i]
	}
	for i:=n-2; i>=0; i--{
		for j:=i+1; j<n; j++{
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[n-1] >= 0
}

func PredictTheWinner_dp(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i:=0; i<n; i++{
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for i:=n-2; i>=0; i--{
		for j:=i+1; j<n; j++{
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

//递归 - 记忆化递归
func PredictTheWinner_recursion_2(nums []int) bool {
	n := len(nums)
	cache := make([][]int, n)
	for i:=0; i<n; i++{
		cache[i] = make([]int, n)
		for j:=0; j<n; j++{
			cache[i][j] = -1 << 63
		}
	}
	return recursion_2(nums, 0, len(nums)-1, cache) >= 0
}

func recursion_2(nums []int, left, right int, cache [][]int)int{
	if left == right{
		return nums[left]
	}
	if cache[left][right] != -1 << 63{
		return cache[left][right]
	}
	pickLeft := nums[left] - recursion_2(nums, left+1, right, cache)
	pickRight := nums[right] - recursion_2(nums, left, right-1, cache)
	max := pickLeft
	if pickLeft < pickRight{
		max = pickRight
	}
	cache[left][right] = max
	return max
}

//递归
func PredictTheWinner_recursion(nums []int) bool {
	return recursion(nums, 0, len(nums)-1) >= 0
}

func recursion(nums []int, left, right int)int{
	if left == right{
		return nums[left]
	}
	pickLeft := nums[left] - recursion(nums, left+1, right)
	pickRight := nums[right] - recursion(nums, left, right-1)
	if pickLeft >= pickRight{
		return pickLeft
	}
	return pickRight
}
