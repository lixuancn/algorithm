package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i >= 1 && nums[i-1] == num {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			if num+nums[left]+nums[right] > 0 {
				right--
			} else if num+nums[left]+nums[right] < 0 {
				left++
			} else {
				res = append(res, []int{num, nums[left], nums[right]})
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
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
}
