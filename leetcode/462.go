package main

import (
	"fmt"
	"sort"
)

//462. 最少移动次数使数组元素相等 II

func minMoves2(nums []int) int {
	//return minMoves2_force(nums)     //暴力法
	//return minMoves2_sort(nums)      //排序 选中间数为目标数，开始移动
	return minMoves2_quickSort(nums) //快速选择 对排序的优化
}

//暴力法 可以通过
func minMoves2_force(nums []int) int {
	result := 1 << 31
	for _, target := range nums {
		ret := 0
		for _, cur := range nums {
			count := target - cur
			if count < 0 {
				count = -count
			}
			ret += count
		}
		if ret < result {
			result = ret
		}
	}
	return result
}

//排序 选中间数为目标数，开始移动
func minMoves2_sort(nums []int) int {
	sort.Ints(nums)
	target := nums[len(nums)>>1]
	result := 0
	for _, num := range nums {
		count := num - target
		if count < 0 {
			count = -count
		}
		result += count
	}
	return result
}

//快速选择 对排序的优化
func minMoves2_quickSort(nums []int) int {
	targetIndex := 0
	var quickSort func(nums []int, left, right int)
	quickSort = func(nums []int, left, right int) {
		if right < left {
			return
		}
		pivot, counter := right, left
		for i := left; i < right; i++ {
			if nums[i] <= nums[pivot] {
				nums[i], nums[counter] = nums[counter], nums[i]
				counter++
			}
		}
		nums[pivot], nums[counter] = nums[counter], nums[pivot]
		k := (len(nums) - 1) >> 1
		if counter == k {
			targetIndex = counter
			return
		} else if counter < k {
			quickSort(nums, counter+1, right)
		} else {
			quickSort(nums, left, counter-1)
		}
	}
	quickSort(nums, 0, len(nums)-1)
	result := 0
	for _, num := range nums {
		count := num - nums[targetIndex]
		if count < 0 {
			count = -count
		}
		result += count
	}
	return result
}

func main() {
	fmt.Println(minMoves2([]int{2, 1, 3})) //2
	//fmt.Println(minMoves2([]int{1, 2, 3}))     //2
	//fmt.Println(minMoves2([]int{1, 10, 2, 9})) //16
}
