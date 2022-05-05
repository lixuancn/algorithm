package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	//return numSubarrayProductLessThanK_recursion(nums, k) // 递归
	return numSubarrayProductLessThanK_doubleIndex(nums, k) // 双指针，滑动窗口
}

//递归
//写完感觉，这不就是两层for的暴力法么，不需要用递归啊
func numSubarrayProductLessThanK_recursion(nums []int, k int) int {
	result := 0
	if k <= 1 {
		return result
	}
	var backtracking func(nums []int, k, start, product int) bool
	backtracking = func(nums []int, k, start, product int) bool {
		if start == len(nums) {
			return false
		}
		product *= nums[start]
		if product >= k {
			return false
		}
		result++
		if !backtracking(nums, k, start+1, product) {
			return false
		}
		return true
	}
	for i := 0; i < len(nums); i++ {
		backtracking(nums, k, i, 1)
	}
	return result
}

// 双指针，滑动窗口
func numSubarrayProductLessThanK_doubleIndex(nums []int, k int) int {
	result := 0
	if k <= 1 {
		return result
	}
	left, right := 0, 0
	product := 1
	for right < len(nums) {
		product *= nums[right]
		//不满足条件，左指针开始右移
		for product >= k {
			product /= nums[left]
			left++
		}
		//加入结果
		//比如left=1,right=1，集合就是1
		//比如left=1,right=2，集合就是2，21
		//比如left=1,right=3，集合就是3，32，321
		//所以[i,j]数组的连续子数组个数为j-i+1
		fmt.Println(right, left)
		result += right - left + 1
		//满足条件，右指针右移
		right++
	}
	return result
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{5, 2, 6}, 100))     //6
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)) //8
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))       //0
}
