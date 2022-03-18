package main

import "fmt"

func jump(nums []int) int {
	end := 0
	max := 0
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		if max < nums[i]+i {
			max = nums[i] + i
		}
		if end == i {
			end = max
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
