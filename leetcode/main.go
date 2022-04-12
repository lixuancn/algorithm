package main

import "fmt"

//一个包括n个元素的有序数组，可能有重复元素。给定t，求满足arr[i] + arr[j] = t (i < j)的 (i, j)对数。
//
//eg：
//-3,-2,-1,0,1,2,2,2,3
//t = 1
//return 5
//
//时间O(n)，

//dp[i]表示放入i之后的对数。 dp[i]
func main() {
	fmt.Println(find([]int{-3, -2, -1, 0, 1, 2, 2, 2, 3}, 5))
}

func find(nums []int, target int) int {
	result := 0
	cache := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		cache[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if _, ok := cache[need]; ok {
			result += cache[need]
		}
	}
	return result
}
