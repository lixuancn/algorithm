package main

import "fmt"

//154. 寻找旋转排序数组中的最小值 II

//自己写的
func findMin(nums []int) int {
	if len(nums) == 1 { //题目保证不为0
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	ans := nums[0]
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] < ans {
			ans = nums[mid]
		}
		if nums[0] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

//官方题解
func findMin_1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] < nums[high] {
			high = mid //这里不能是mid-1,比如[4,5,1,2,3]，mid=1，right=mid-1会错过真正的最小值
		} else {
			low = mid + 1 //这里可以mid+1，因为mid>=right会进入else分支，那么最小值就肯定不是mid了
		}
	}
	return nums[low] //因为上面low=mid+1，mid>=right那么mid肯定不是最小值，mid+1有可能是，low已经是mid+1了。
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
