package main

import "fmt"

//189. 轮转数组

func rotate(nums []int, k int) {
	//rotate_extendArr(nums, k) //额外数组
	//rotate_force(nums, k) //暴力法
	rotate_reverse(nums, k) //数组翻转
}

//使用额外数组
func rotate_extendArr(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	arr := make([]int, n)
	for i := range nums {
		arr[(i+k)%n] = nums[i]
	}
	copy(nums, arr)
}

//暴力法
func rotate_force(nums []int, k int) {
	n := len(nums)
	k = k % n
	for m := 0; m < k; m++ {
		for i := 0; i < n; i++ {
			nums[i], nums[n-1] = nums[n-1], nums[i]
		}
	}
}

//翻转数组
func rotate_reverse(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	reverse := func(nums []int, left, right int) {
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums := []int{1, 2}
	rotate(nums, 1)
	fmt.Println(nums)
}
