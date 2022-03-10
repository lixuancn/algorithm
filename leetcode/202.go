package main

import (
	"fmt"
)

func isHappy(n int) bool {
	record := make(map[int]struct{})
	for {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
		if sum == 1 {
			return true
		}
		if _, ok := record[sum]; ok {
			return false
		}
		record[sum] = struct{}{}
	}
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
