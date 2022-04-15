package main

import (
	"fmt"
)

//7. 整数反转

func reverse(x int) int {
	result := 0
	for x != 0 {
		if -1<<31/10 > result || 1<<31/10 < result {
			return 0
		}
		v := x % 10
		x = x / 10
		result = result*10 + v
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}
