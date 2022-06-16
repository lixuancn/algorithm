package main

import (
	"fmt"
)

//1021. 删除最外层的括号

//栈改为计数，可以优化空间
func removeOuterParentheses(s string) string {
	count := 0
	start := 0
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else {
			count--
		}
		if count == 0 {
			result += s[start : i+1][1 : i-start]
			start = i + 1
		}
	}
	return result
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
