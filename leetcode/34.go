package main

import (
	"fmt"
)

//34. 在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			result[0] = mid
			result[1] = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if result[0] == -1 {
		return result
	}
	left, right = result[0]+1, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			result[1] = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	//fmt.Println(searchRange([]int{8, 8}, 8)) //3,4
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) //3,4
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) //-1,-1
	fmt.Println(searchRange([]int{5, 7, 7, 8, 10}, 8))    //3,3
	fmt.Println(searchRange([]int{}, 0))                  //-1,-1
}
