package main

import "fmt"

//66. 加一

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}
