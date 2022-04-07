package main

import "fmt"

//152. 乘积最大子数组

func maxProduct(nums []int) int {
	//dp[i]表示前i个元素组成的子数组中，最大乘积。但是因为有正负，所以维护两个dp，一个表示最大值越大越好，一个表示最小值，越小越好（负数，希望以后可以变正）
	//dpMax[i] = max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
	//dpMin[i] = min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	result := dpMax[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = max(dpMax[i-1]*nums[i], max(dpMin[i-1]*nums[i], nums[i]))
		dpMin[i] = min(dpMax[i-1]*nums[i], min(dpMin[i-1]*nums[i], nums[i]))
		if dpMax[i] > result {
			result = dpMax[i]
		}
	}
	return result
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
	fmt.Println(maxProduct([]int{2, 3, -2, 4})) //6
	fmt.Println(maxProduct([]int{-2, 0, -1}))   //0
}
