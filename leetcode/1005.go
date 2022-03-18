package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		if nums[i] == 0 {
			k = 0
			break
		}
	}
	if k == 0 || k%2 == 0 {
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	sort.Ints(nums)
	nums[0] = -nums[0]
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
}
