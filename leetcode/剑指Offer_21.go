package main

import "fmt"

//我自己想的 - 奇数顺序不变
func exchange(nums []int) []int {
	left, right := 0, 1
	for right < len(nums) && left < len(nums) {
		//如果是偶数
		if nums[left]&1 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right++
			if nums[left]&1 != 0 {
				left++
			}
		} else {
			left++
			right++
		}
	}
	return nums
}

//题解的双指针 - 奇数顺序不变
func exchange_1(nums []int) []int {
	slow, fast := 0, 0
	for fast < len(nums) {
		//快指针是奇数
		if nums[fast]&1 == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}

//双指针 - 奇数顺序汇编
func exchange_2(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]&1 != 0 {
			left++
			continue
		}
		if nums[right]&1 != 1 {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

func main() {
	nums := []int{1, 2, 4, 5, 1, 6, 8, 13}
	exchange(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 4, 5, 1, 6, 8, 13}
	exchange_1(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 4, 5, 1, 6, 8, 13}
	exchange_2(nums)
	fmt.Println(nums)
}
