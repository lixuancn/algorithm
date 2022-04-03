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

func sortArrayByParityII_practice(nums []int) []int {
	//i表示偶数，j表示奇数
	for i, j := 0, 1; i < len(nums) && j < len(nums); i, j = i+2, j+2 {
		//i找奇数
		for ; i < len(nums) && nums[i]&1 == 0; i += 2 {
		}
		//j找偶数
		for ; j < len(nums) && nums[j]&1 == 1; j += 2 {
		}
		//交换
		if i < len(nums) && j < len(nums) && nums[i]&1 == 1 && nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParityII_practice([]int{4, 2, 5, 7}))
}
