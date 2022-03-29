package main

import "fmt"

//二分的巧妙方法
//如果这个数是x，那么x左边和右边都是成对出现，左边，如果nums[y]=nums[y+1]，则y是偶数。右边，左边，如果nums[y]=nums[y+1]，则y是奇数。
//二分看mid是偶数还是奇数
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		//这里就不判断奇偶了，利用异或的特性。
		//当i是偶数时，mid + 1 = mid ^ 1
		//当i是奇数时，mid - 1 = mid ^ 1
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

//这个方法属于递归了。
func singleNonDuplicate_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	} else if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	return find(nums, 1, len(nums)-2)
}

func find(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (right-left)/2 + left
	if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
		return nums[mid]
	}
	r := find(nums, left, mid-1)
	if r != -1 {
		return r
	}
	r = find(nums, mid+1, right)
	return r
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
