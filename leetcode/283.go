package main

import "fmt"

//283. 移动零

func moveZeroes(nums []int) {
	left, right := 0, 0
	for left <= right && right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
