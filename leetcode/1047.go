package main

import "fmt"

func removeDuplicates(s string) string {
	stack := make([]byte, 1)
	stack[0] = s[0]
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
