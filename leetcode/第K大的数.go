package main

import "fmt"

//最大的第K个数

func main() {
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 9))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 8))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 7))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 6))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 5))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 4))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 3))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 2))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 1))
}

func quickSelect(nums []int, start, end, k int) int {
	pivot := end
	counter := start
	for i := start; i < end; i++ {
		//正序
		if nums[i] <= nums[pivot] {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	if end-counter+1 == k {
		return nums[counter]
	}
	if end-counter+1 > k {
		//取右边
		return quickSelect(nums, counter+1, end, k)
	} else {
		//取左边
		return quickSelect(nums, start, counter-1, k-(end-counter+1))
	}
}
