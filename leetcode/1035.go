package main

//1035. 不相交的线

//在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
//1、nums1[i] == nums2[j]
//2、且绘制的直线不与任何其他连线（非水平线）相交。
//请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//以这种方法绘制线条，并返回可以绘制的最大连线数。

//最大子串的问题
import (
	"fmt"
)

func main() {
	p1 := []int{2, 5, 1, 2, 5}
	p2 := []int{10, 5, 2, 1, 5, 2}

	p1 = []int{1, 2, 4}
	p2 = []int{2, 1, 4}

	p1 = []int{1, 3, 7, 1, 7, 5}
	p2 = []int{1, 9, 2, 5, 1}
	r := maxUncrossedLines(p1, p2)
	fmt.Println(r)
}

//动态规划
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//新来的两个元素可以连线
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

//根据题意，相等的数字可以连线，连线不能相交。那么就转换为：有相同的数字，且顺序不变。
//那么这道题可以转换为求最长公共子序列的问题。
func maxUncrossedLines_dp_practice(nums1 []int, nums2 []int) int {
	//dp[i][j]表示[0,i-1]和[0,j-1]的两个子数组的最长公共子序列
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
