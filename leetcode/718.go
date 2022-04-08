package main

import "fmt"

//718. 最长重复子数组

func findLength(nums1 []int, nums2 []int) int {
	//return findLength_dp_1(nums1, nums2) //动态规划
	//return findLength_dp_2(nums1, nums2) //动态规划
	return findLength_window(nums1, nums2) //滑动窗口
}

//动态规划
//dp[i][j]表示[0,i]和[0,j]两个数组的最长重复数组
//nums[i] == nums[j]，则dp[i][j] = dp[i-1][j-1]+1
//因为循环依赖i-1，所以dp[i][j]表示数字nums[i-1]和nums[j-1]这样会比较方便，也就是循环从(1,n]
//如果从[0,n)这样循环的话，初始化会比较麻烦，需要分别堆dp[0][j]和dp[i][0]都初始化了
//函数_dp1和_dp2分别对应的是不同的初始化
func findLength_dp_1(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2))
	}
	result := 0
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
			result = 1
		}
	}
	for j := 0; j < len(nums2); j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
			result = 1
		}
	}
	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if result < dp[i][j] {
				result = dp[i][j]
			}
		}
	}
	return result
}

func findLength_dp_2(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	result := 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if result < dp[i][j] {
				result = dp[i][j]
			}
		}
	}
	return result
}

//滑动窗口，想象成两把尺子，相向而行 https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/wu-li-jie-fa-by-stg-2/
//todo 没看懂题解
func findLength_window(nums1 []int, nums2 []int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n, m := len(nums1), len(nums2)
	result := 0
	for i := 1; i <= n; i++ {
		result = max(result, maxLen(nums1, 0, nums2, m-i, i))
	}
	for j := m - n; j >= 0; j-- {
		result = max(result, maxLen(nums1, 0, nums2, j, n))
	}
	for i := 1; i < n; i++ {
		result = max(result, maxLen(nums1, i, nums2, 0, n-i))
	}
	return result
}

func maxLen(nums1 []int, i int, nums2 []int, j int, length int) int {
	count, result := 0, 0
	for k := 0; k < length; k++ {
		if nums1[i+k] == nums2[j+k] {
			count++
		} else {
			if result < count {
				result = count
			}
			count = 0
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

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 8}, []int{5, 6, 1, 4, 7})) //1
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})) //3
	fmt.Println(findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0})) //5
}
