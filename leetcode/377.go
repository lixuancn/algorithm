package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 0
	//求排列，所以背包在外层循环，物品在内层循环，可看518题的详细解释
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}

func main() {
	combinationSum4([]int{1, 2, 3}, 4)
}
