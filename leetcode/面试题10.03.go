package main

import "fmt"

func search(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] == target {
			if mid == 0 {
				return 0
			} else if mid > 0 && arr[mid] != arr[mid-1] {
				return mid
			}
		} else if arr[left] == arr[mid] {

		} else if arr[left] <= arr[mid] {
			
		}
	}
}

func main() {
	fmt.Println(search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5))  //8
	fmt.Println(search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 11)) //-1
	fmt.Println(search([]int{1}, 11))                                         //-1
}
