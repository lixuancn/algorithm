package main

import "fmt"

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] { //这里我觉得len(nums)-1可以换成right。上面的0同理可以换成left，不过不影响结果。
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func search_practice(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right && left < n && right >= 0 {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search_practice([]int{4, 5, 6, 7, 0, 1, 2}, 0)) //4
	fmt.Println(search_practice([]int{1}, 1))                   //0
	fmt.Println(search_practice([]int{1, 3}, 1))                //0
	fmt.Println(search_practice([]int{1, 3}, 3))                //1
	fmt.Println(search_practice([]int{3, 1}, 1))                //1
	fmt.Println(search_practice([]int{1, 3, 5}, 1))             //0
	fmt.Println(search_practice([]int{5, 1, 3}, 5))             //0
}
