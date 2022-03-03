package main

import "fmt"

//快慢指针
func removeDuplicates(nums []int) int {
	slow, fast := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	var nums []int
	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
