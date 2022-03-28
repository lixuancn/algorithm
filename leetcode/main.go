package main

import "fmt"

//1,2,4,5,6,8
//2,8,34,56
//输出：
//1,2,4,5,6,8,34,56
func main() {
	nums1 := []int{1, 2, 4, 5, 6, 8}
	nums2 := []int{2, 8, 34, 56}
	nums := merge(nums1, nums2)
	fmt.Println(nums)
}

func merge(nums1, nums2 []int) []int {
	nums := make([]int, 0)
	index1, index2 := 0, 0
	for index1 < len(nums1) && index2 < len(nums2) {
		num := nums1[index1]
		if nums1[index1] > nums2[index2] {
			num = nums2[index2]
			index2++
		} else {
			index1++
		}
		//已有过滤
		if len(nums) > 0 && num == nums[len(nums)-1] {
			continue
		}
		nums = append(nums, num)
	}
	//2合并完了，1和没合并完
	for i := index1; i < len(nums1); i++ {
		if len(nums) > 0 && nums1[i] == nums[len(nums)-1] {
			continue
		}
		nums = append(nums, nums1[i])
	}
	//1合并完了，2和没合并完
	for i := index2; i < len(nums2); i++ {
		if len(nums) > 0 && nums2[i] == nums[len(nums)-1] {
			continue
		}
		nums = append(nums, nums2[i])
	}
	return nums
}
