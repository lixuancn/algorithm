package main

import "fmt"

//977. 有序数组的平方

//双指针
func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	ret := make([]int, n, n)
	i := n - 1
	for ; left <= right; i-- {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			ret[i] = nums[left] * nums[left]
			left++
		} else {
			ret[i] = nums[right] * nums[right]
			right--
		}
	}
	return ret
}

//找到正负数的分界点，那么平方后，左边是单调递减，右边是单调递增。可以用类似归并排序的方式来对两个有序数组排序
func sortedSquares_2(nums []int) []int {
	ans := make([]int, len(nums))
	ansIndex := 0
	lastNegIndex := -1
	for i := 0; i < len(nums) && nums[i] < 0; i++ {
		lastNegIndex = i
	}
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < len(nums); {
		//如果j走完了，
		if j == len(nums) {
			ans[ansIndex] = nums[i] * nums[i]
			i--
		} else if i < 0 {
			ans[ansIndex] = nums[j] * nums[j]
			j++
		} else if nums[i]*nums[i] >= nums[j]*nums[j] {
			ans[ansIndex] = nums[j] * nums[j]
			j++
		} else {
			ans[ansIndex] = nums[i] * nums[i]
			i--
		}
		ansIndex++
	}
	return ans
}

func main() {
	fmt.Println(sortedSquares_2([]int{-4, -1, 0, 3, 10}))
	//fmt.Println(sortedSquares_2([]int{-7, -3, 2, 3, 11}))
	//fmt.Println(sortedSquares_2([]int{-5, -3, -2, -1}))
}
