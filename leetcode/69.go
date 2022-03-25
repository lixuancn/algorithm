package main

import "fmt"

//平方根

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 0, x
	ans := 0
	for left <= right {
		mid := (right-left)/2 + left
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
