package main

import (
	"fmt"
	"strconv"
)

//241. 为运算表达式设计优先级

func diffWaysToCompute(expression string) []int {
	num, err := strconv.Atoi(expression)
	if err == nil {
		return []int{num}
	}
	result := []int{}
	for i, char := range expression {
		s := string(char)
		if s == "+" || s == "-" || s == "*" {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, leftNum := range left {
				for _, rightNum := range right {
					ret := 0
					if s == "+" {
						ret = leftNum + rightNum
					} else if s == "-" {
						ret = leftNum - rightNum
					} else {
						ret = leftNum * rightNum
					}
					result = append(result, ret)
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}
