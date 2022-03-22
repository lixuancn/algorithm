package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	//guibingpaixu(nums1, m, nums2, n) //归并排序
	//doublePointer_1(nums1, m, nums2, n) //利用两个数组已经有序，双指针 正向 需要空间
	doublePointer_2(nums1, m, nums2, n) //利用两个数组已经有序，双指针 反向 不需要空间，利用nums1后半部分为空的特性
}

func doublePointer_1(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, 0, len(nums1))
	index1, index2 := 0, 0
	for {
		if index1 == m {
			tmp = append(tmp, nums2[index2:]...)
			break
		}
		if index2 == n {
			tmp = append(tmp, nums1[index1:]...)
			break
		}
		if nums1[index1] <= nums2[index2] {
			tmp = append(tmp, nums1[index1])
			index1++
		} else {
			tmp = append(tmp, nums2[index2])
			index2++
		}
	}
	copy(nums1, tmp)
}

func doublePointer_2(nums1 []int, m int, nums2 []int, n int) {
	index1, index2, tailIndex := m-1, n-1, len(nums1)-1
	for index1 >= 0 || index2 >= 0 {
		if index1 == -1 {
			nums1[tailIndex] = nums2[index2]
			index2--
		} else if index2 == -1 {
			nums1[tailIndex] = nums1[index1]
			index1--
		} else if nums1[index1] > nums2[index2] {
			nums1[tailIndex] = nums1[index1]
			index1--
		} else {
			nums1[tailIndex] = nums2[index2]
			index2--
		}
		tailIndex--
	}
}

func guibingpaixu(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}
	mergeSort(nums1, 0, len(nums1)-1)
}

//归并排序 模板
func mergeSort(nums []int, left, right int) {
	if left == right {
		return
	}
	mid := (right-left)/2 + left
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	mergeResult(nums, left, mid, right)
}

func mergeResult(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j, tmpIndex := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[tmpIndex] = nums[i]
			i++
		} else {
			tmp[tmpIndex] = nums[j]
			j++
		}
		tmpIndex++
	}
	for i <= mid {
		tmp[tmpIndex] = nums[i]
		i++
		tmpIndex++
	}
	for j <= right {
		tmp[tmpIndex] = nums[j]
		j++
		tmpIndex++
	}
	for p := 0; p < len(tmp); p++ {
		nums[left+p] = tmp[p]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	//nums1 := []int{2, 0}
	//merge(nums1, 1, []int{1}, 1)
	fmt.Println(nums1)
}
