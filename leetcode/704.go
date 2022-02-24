package main

import "fmt"

func search(nums []int, target int) int {
	left, right, middle := 0, len(nums)-1, 0
	for left <= right{
		middle = (right - left) / 2 + left
		if target == nums[middle]{
			return middle
		}
		if target < nums[middle]{
			right = middle - 1
		}else{
			left = middle + 1
		}
	}
	return -1
}

//如果while的条件是<
func search2(nums []int, target int) int {
	left, right, middle := 0, len(nums)-1, 0
	for left < right{
		middle = (right - left) / 2 + left
		if target == nums[middle]{
			return middle
		}
		if target < nums[middle]{
			right = middle
		}else{
			left = middle + 1
		}
	}
	return -1
}

func main(){
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9))
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
}