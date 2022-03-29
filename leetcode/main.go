package main

import "fmt"

//[3,5,5,5,6,6,6,7,7,9,9,12]
var min, max = 2 << 31, -1

func main() {
	nums := []int{3, 5, 5, 5, 6, 6, 6, 7, 7, 9, 9, 12}
	find(nums, 6, 0, len(nums))
	fmt.Println(max - min)

	find(nums, 3, 0, len(nums))
	find(nums, 5, 0, len(nums))
	find(nums, 6, 0, len(nums))
	find(nums, 7, 0, len(nums))
	find(nums, 9, 0, len(nums))
	find(nums, 12, 0, len(nums))
}

func find(nums []int, target, left, right int) {
	if left >= right {
		if left > max {
			max = left
		}
		if left < min {
			min = left
		}
		return
	}
	mid := (right-left)/2 + left
	//在右边
	if nums[mid] < target {
		find(nums, target, mid+1, right)
		//在左边
	} else if nums[mid] > target {
		find(nums, target, left, mid)
	} else {
		find(nums, target, left, mid)
		find(nums, target, mid+1, right)
	}
}
