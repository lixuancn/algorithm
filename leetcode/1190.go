package main
import (
	"fmt"
)

func main() {
	r := reverseParentheses("(u(love)i)")
	fmt.Println(r)
}

func reverseParentheses(s string) string {
	//先找出所有的括号对
	n := len(s)
	pair := make([]int, n)
	stack := make([]int, 0)
	for i := range s{
		if s[i] == '('{
			stack = append(stack, i)
		}else if s[i] == ')'{
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[left], pair[i] = i, left
		}
	}
	ret := []byte{}
	for i,step:=0,1; i<n; i+=step{
		if s[i] == '(' || s[i] == ')'{
			i = pair[i]
			step = -step
		}else{
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}

func reverseParentheses_stack_loop(s string) string {
	stack := make([][]byte, 0)
	str := make([]byte, 0)
	for i := range s{
		if s[i] == '('{
			stack = append(stack, str)
			str = make([]byte, 0)
		}else if s[i] == ')'{
			for j,n:=0,len(str); j<n/2; j++{
				str[j], str[n-1-j] = str[n-1-j], str[j]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		}else{
			str = append(str, s[i])
		}
	}
	return string(str)
}
