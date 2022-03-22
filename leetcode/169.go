package main

import "fmt"

func majorityElement(nums []int) int {
	//return majorityElement_moore(nums) //摩尔投票，是最优解
	return majorityElement_recursion(nums) //分治
}

//摩尔投票法
//默认第一个数是一拨人，其他数是一拨人，循环是如果相同，就说明自己人，人数加一，如果不同，就自杀式干掉对方。
//最后count >= 1时说明是多数元素
func majorityElement_moore(nums []int) int {
	count := 0
	var ret int
	for i, num := range nums {
		if count == 0 {
			ret = nums[i]
		}
		if num == ret {
			count++
		} else {
			count--
		}
	}
	return ret
}

//nums一分为二，nums的众数至少在分割后的某一部分里也是众数
func majorityElement_recursion(nums []int) int {
	return recursion(nums, 0, len(nums)-1)
}

func recursion(nums []int, low, height int) int {
	//数组长度为1的数组，唯一哪个元素就是众数
	if low == height {
		return nums[low]
	}
	mid := (height-low)/2 + low
	left := recursion(nums, low, mid)
	right := recursion(nums, mid+1, height)
	if left == right {
		return left
	}
	leftCount := countInRange(nums, left, low, height)
	rightCount := countInRange(nums, right, low, height)
	if leftCount > rightCount {
		return left
	}
	return right
}

func countInRange(nums []int, target, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3, 1, 0}))
}
