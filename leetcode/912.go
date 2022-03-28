package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	sortArray_quickSort(nums, 0, len(nums)-1) //快排
	//sortArray_merge(nums, 0, len(nums)-1) //归并
	//sortArray_heap(nums) //堆
	return nums
}

func sortArray_quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot, counter := end, start
	for i := start; i < end; i++ {
		if nums[i] <= nums[pivot] {
			nums[i], nums[counter] = nums[counter], nums[i]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	sortArray_quickSort(nums, start, counter-1)
	sortArray_quickSort(nums, counter+1, end)
}

func sortArray_merge(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (right + left) >> 1
	sortArray_merge(nums, left, mid)
	sortArray_merge(nums, mid+1, right)
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for ; i <= mid; i++ {
		tmp[k] = nums[i]
		k++
	}
	for ; j <= right; j++ {
		tmp[k] = nums[j]
		k++
	}
	for i = 0; i < len(tmp); i++ {
		nums[left+i] = tmp[i]
	}
}

func sortArray_heap(nums []int) []int {
	n := len(nums) - 1
	//1、构建一个最大堆
	for i := n / 2; i >= 0; i-- {
		maxHeapify(nums, i, n)
	}
	for i := n; i >= 1; i-- {
		//2、把一个数放到堆顶
		nums[i], nums[0] = nums[0], nums[i]
		n--
		//3、下移
		maxHeapify(nums, 0, n)
	}
	return nums
}

func maxHeapify(nums []int, parent, n int) {
	for 2*parent+1 <= n {
		leftChild := 2*parent + 1
		rightChild := 2*parent + 2
		//取最大的节点
		largeNode := parent
		if leftChild <= n && nums[leftChild] > nums[largeNode] {
			largeNode = leftChild
		}
		if rightChild <= n && nums[rightChild] > nums[largeNode] {
			largeNode = rightChild
		}
		//最大的不是父节点，就和最大子节点换位置
		if largeNode != parent {
			nums[parent], nums[largeNode] = nums[largeNode], nums[parent]
			parent = largeNode
		} else {
			break
		}
	}
}

func main() {
	//fmt.Println(sortArray([]int{2, 5, 7, 9, 1, 3, 4, 8, 6}))
	fmt.Println(sortArray([]int{0, 0, 1, 2, 5, 1}))

}
