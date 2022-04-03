package main

import "fmt"

//209. 长度最小的子数组

//滑动窗口
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, right := 0, 0
	sum := 0
	size := 0
	for ; right < n; right++ {
		sum += nums[right]
		for left <= right && sum >= target {
			if right-left+1 < size || size == 0 {
				size = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	return size
}

func main() {
	var nums []int
	nums = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
	nums = []int{1, 4, 4}
	fmt.Println(minSubArrayLen(4, nums))
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(11, nums))
}
