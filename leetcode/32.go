package main

import "fmt"

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	//dp[i]表示以s[i]结尾的有效子串的长度，所以如果dp[i]=='('，dp[i]=0
	dp := make([]int, len(s))
	result := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				//我是一个），我前一个还是），那就要之前还有没有没用到的）
				if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
					if i-1-dp[i-1]-1 >= 0 {
						dp[i] = dp[i-1] + 2 + dp[i-1-dp[i-1]-1]
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func main() {
	fmt.Println(longestValidParentheses("(()"))    //2
	fmt.Println(longestValidParentheses(")()())")) //4
	fmt.Println(longestValidParentheses(""))       //0
	fmt.Println(longestValidParentheses("()"))     //0
}
