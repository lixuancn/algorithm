package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	stack := make([]byte, 0)
	if len(s) == 0 {
		return true
	}
	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != byte(c) {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid_practice(s string) bool {
	stack := list.New()
	for _, c := range s {
		if c == '(' {
			stack.PushBack(')')
		} else if c == '[' {
			stack.PushBack(']')
		} else if c == '{' {
			stack.PushBack('}')
		} else if stack.Len() == 0 {
			return false
		} else if stack.Remove(stack.Back()).(rune) != c {
			return false
		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid_practice("()"))
}
