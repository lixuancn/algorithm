package main

import (
	"fmt"
)

//面试题 08.09. 括号
var ret []string

func generateParenthesis(n int) []string {
	ret = make([]string, 0)
	backtracking("", n, n)
	return ret
}

func backtracking(str string, left, right int) {
	if left == 0 && right == 0 {
		ret = append(ret, str)
		return
	}
	if left > 0 {
		backtracking(str+"(", left-1, right)
	}
	//这里，不是right>0，而是str中右括号数量要小于左括号的时候才可以加
	if right > left {
		backtracking(str+")", left, right-1)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
