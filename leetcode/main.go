package main

import (
	"fmt"
)

type T struct {
	a int
	b string
}

func main() {
	fmt.Println(match("([()])[]{}"))
}

func match(s string) bool {
	stack := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else {
			//取栈第一个元素
			c := stack[len(stack)-1]
			//pop
			stack = stack[:len(stack)-1]
			if c == s[i] {
				continue
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
