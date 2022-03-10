package main

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
