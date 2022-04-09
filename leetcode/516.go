package main

import "fmt"

//516. 最长回文子序列

func longestPalindromeSubseq(s string) int {
	//dp[i][j]表示[i,j]子字符串中的最长回文子序列
	//s[i]==s[j]，那么就是[i+1,j-1]中的最长回文子序列+2。dp[i][j] = dp[i-1][j-1]+2
	//s[i]!=s[j]，那么就看[i+1,j]和[i，j-1]中取大的。dp[i][j] = max(dp[i+1][j], dp[i][j-1])
	//初始化时，单个字符也是回文字符串
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[0][len(s)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(longestPalindromeSubseq("cbbd"))  //2
	fmt.Println(longestPalindromeSubseq("bbbab")) //4
}
