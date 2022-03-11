package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			v1 := stack[len(stack)-1]
			v2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, v2+v1)
			case "-":
				stack = append(stack, v2-v1)
			case "*":
				stack = append(stack, v2*v1)
			case "/":
				stack = append(stack, v2/v1)
			}
		}
	}
	return stack[len(stack)-1]
}

func main() {
	//fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
