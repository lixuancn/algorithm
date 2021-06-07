package main
import (
	"fmt"
)

func main() {
	r := findTargetSumWays([]int{1,1,1,1,1}, 3)
	fmt.Println(r)
}

//动态规划
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num :=range nums{
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff % 2 == 1{
		return 0
	}
	n, neg := len(nums), diff / 2
	dp := make([][]int, n+1)
	for i := range dp{
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums{
		for j:=0; j<=neg; j++{
			dp[i+1][j] = dp[i][j]
			if j >= num{
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}

//递归
var count int
func findTargetSumWays_recursion(nums []int, target int) int {
	count = 0
	recursion(nums, target, len(nums), 0, nums[0])
	recursion(nums, target, len(nums), 0, -nums[0])
	return count
}

func recursion(nums []int, target, n, index, sum int){
	if index == n - 1{
		if sum == target{
			count++
		}
		return
	}
	i := index + 1
	recursion(nums, target, n, i, sum+nums[i])
	recursion(nums, target, n, i, sum-nums[i])
}
