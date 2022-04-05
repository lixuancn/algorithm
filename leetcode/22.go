package main

import "fmt"

//22. 括号生成

var result []string

func generateParenthesis(n int) []string {
	result = []string{}
	if n == 0 {
		return result
	}
	backtracking(n, "", 0, 0)
	return result
}

func backtracking(n int, str string, leftCount, rightCount int) {
	if leftCount == n && rightCount == n {
		result = append(result, str)
		return
	}
	if leftCount < n {
		backtracking(n, str+"(", leftCount+1, rightCount)
	}
	if rightCount < leftCount {
		backtracking(n, str+")", leftCount, rightCount+1)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
