package main

import "fmt"

//647. 回文子串

func countSubstrings(s string) int {
	return countSubstrings_dp(s)
	return countSubstrings_doublePointer(s)
}

//dp[i][j]表示[i,j]之间是不是回文串
//s[i]!=s[j]，则dp[i][j] = false
//s[i]==s[j]，如果i==j说明只有一个字符，如a，则是回文串
//s[i]==s[j]，如果j-i==1说明只有两个字符，如aa，则是回文串
//s[i]==s[j]，如果j-i>1说明有多于两个字符，如cabac，c相等，则需要判断，[i+1,j-1]是不是回文串（aba）
//s[i]==s[j]，dp[i][j] = dp[i+1][j-1]。由公式可以看出，i需要倒退，j需要正推。j从i开始，即从中间向两边开始扫描s
func countSubstrings_dp(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					result++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	return result
}

//从中心点向两边扩散，中心点有两种情况，一个中心点和两个中心点
func countSubstrings_doublePointer(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result += scan(s, i, i)   //以i为中心
		result += scan(s, i, i+1) //以i，i+1为中心
	}
	return result
}

func scan(s string, i, j int) int {
	result := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		result++
	}
	return result
}

func main() {
	fmt.Println(longestPalindromeSubseq("abc"))
	//fmt.Println(longestPalindromeSubseq("aaa"))
}
