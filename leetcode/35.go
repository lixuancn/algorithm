package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right, middle := 0, n-1, 0
	ret := n
	for left <= right{
		middle = (right - left) >> 1 + left
		if target <= nums[middle]{
			ret = middle
			right = middle - 1
		}else{
			left = middle + 1
		}
	}
	return ret
}

func main(){
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
	fmt.Println(searchInsert([]int{1,3,5,6}, 0))
	fmt.Println(searchInsert([]int{1}, 0))
}