package main

import "fmt"

//282. 给表达式添加运算符

func addOperators(num string, target int) []string {
	n := len(num)
	result := make([]string, 0)
	var backtracking func(expr []byte, i, res, mul int)
	backtracking = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				result = append(result, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)
		}
		//nums[j] != 0的意思是，可以是单个0，但不能是前导0
		for j, val := i, 0; j < n && (i == j || num[j] != 0); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				backtracking(expr, j+1, val, val) //开头不能是符号开头
			} else {
				expr[signIndex] = '+'
				backtracking(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtracking(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				backtracking(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtracking(make([]byte, 0, n*2-1), 0, 0, 0)
	return result
}

func main() {
	//fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("232", 8))
}
