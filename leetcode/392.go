package main

import "fmt"

//392. 判断子序列

func isSubsequence(s string, t string) bool {
	return isSubsequence_doublePointer(s, t) //双指针的解法
	return isSubsequence_dp(s, t)            //动态规划
}

func isSubsequence_doublePointer(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

//也可以转换为求子序列的长度。如果子序列长度==s的长度，则返回true
func isSubsequence_dp(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				//如果没找到，相当于t的指针要继续往后走，s的指针不动
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
