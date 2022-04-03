package main

//209. 长度最小的子数组

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
