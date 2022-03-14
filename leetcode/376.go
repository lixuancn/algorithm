package main

import "fmt"

//贪心
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	ret := 1
	preDiff := nums[1] - nums[0]
	if preDiff != 0 {
		ret = 2
	}
	for i := 2; i < n; i++ {
		curDiff := nums[i] - nums[i-1]
		if preDiff >= 0 && curDiff < 0 || preDiff <= 0 && curDiff > 0 {
			ret++
			preDiff = curDiff
		}
	}
	return ret
}

//动态规划
func wiggleMaxLength_dp(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(up[i-1]+1, down[i-1])
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	//fmt.Println(wiggleMaxLength([]int{1, 2, 1, 1, 1}))
	//fmt.Println(wiggleMaxLength([]int{1, 1, 1, 1, 1}))
	//fmt.Println(wiggleMaxLength([]int{9, 9, 9}))
	//fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	//fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}
