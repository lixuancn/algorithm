package main

import "fmt"

//81. 搜索旋转排序数组 II
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return nums[0] == target
	}
	left, right := 0, len(nums)-1
	for left <= right {
		fmt.Println(left, right)
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[left] == nums[mid] && nums[mid] == nums[right] {
			//比如[3，1，2，3，3，3,3],target=2 第一次循环时，无法判断target属于[0,3]还是[4,6]的区间
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

func main() {
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	//fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
}
