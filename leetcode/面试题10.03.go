package main

import "fmt"

//面试题 10.03. 搜索旋转数组

func search(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		//返回最左的
		if arr[left] == target {
			return left
		}
		mid := (left + right) >> 1
		if arr[mid] == target {
			right = mid //因为要返回下标最小的，所以right直接移到mid来
		} else if arr[left] == arr[mid] {
			left++
		} else if arr[left] < arr[mid] { //左边是有序的
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5))  //8
	fmt.Println(search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 11)) //-1
	fmt.Println(search([]int{1}, 11))                                         //-1
	fmt.Println(search([]int{1, -2}, -2))                                     //1
}
