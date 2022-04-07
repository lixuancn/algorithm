package main

import (
	"fmt"
)

//53. 最大子数组和

//贪心
//如果加和后是负数，就是抛弃不要，从当前节点重新开始加。因为负数会越加越小
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	ret := -1 << 31
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > ret {
			ret = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return ret
}

//动态规划
func maxSubArray_dp(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSubArray_tanxin_practice(nums []int) int {
	result := 0
	sum, result := 0, -1<<31
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > result {
			result = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return result
}

func maxSubArray_dp_practice(nums []int) int {
	//dp[i]表示长度为i的子数组的最大子数组的和
	//dp[i] = max(dp[i-1], dp[i]+nums[i])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func main() {
	fmt.Println(maxSubArray_dp_practice([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println(maxSubArray_dp_practice([]int{1}))                             //1
	fmt.Println(maxSubArray_dp_practice([]int{-1, -1, -1, -1}))                //-1
	fmt.Println(maxSubArray_dp_practice([]int{5, 4, -1, 7, 8}))                //23
}
