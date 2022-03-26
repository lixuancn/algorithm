package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	ou, ji := 0, 1
	for ; ou < len(nums); ou += 2 {
		if nums[ou]&1 == 1 {
			for ji < len(nums) && nums[ji]&1 == 1 {
				ji += 2
			}
			nums[ou], nums[ji] = nums[ji], nums[ou]
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}
