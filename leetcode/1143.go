package main

import "fmt"

//1143. 最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j]表示子序列[0,i-1]和[0,j-1]的最长公共子序列
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace")) //3
	fmt.Println(longestCommonSubsequence("abc", "abc"))   //3
	fmt.Println(longestCommonSubsequence("abc", "def"))   //0
}
