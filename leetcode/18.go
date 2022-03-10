package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if i >= 1 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}

func main() {
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}
