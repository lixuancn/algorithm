package main

import "fmt"

func main() {
	//fmt.Println(maopao([]int{7, 3, 1, 2, 9, 4, 8, 6, 5})) //冒泡
	//fmt.Println(kuaipai([]int{7, 3, 1, 2, 9, 4, 8, 6, 5})) //快排
	//fmt.Println(guibing([]int{7, 3, 1, 2, 9, 4, 8, 6, 5})) //归并
	fmt.Println(duipaixu([]int{7, 3, 1, 2, 9, 4, 8, 6, 5})) //堆排序
}

//冒泡
func maopao(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j <
			len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

//快排
func kuaipai(nums []int) []int {
	var recursion func(nums []int, left, right int)
	recursion = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		//pivot: 标杆的下标。  counter：比pivot小的元素
		pivot, counter := right, left
		//比标杆小的都放在左边
		for i := left; i < right; i++ {
			if nums[i] < nums[pivot] {
				nums[i], nums[counter] = nums[counter], nums[i]
				counter++
			}
		}
		//该挪动标杆了
		nums[pivot], nums[counter] = nums[counter], nums[pivot]
		//递归
		recursion(nums, left, pivot-1)
		recursion(nums, pivot+1, right)
	}
	recursion(nums, 0, len(nums)-1)
	return nums
}

//归并
func guibing(nums []int) []int {
	var mergeSort func(nums []int, left, right int)
	var merge func(nums []int, left, mid, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)/2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
	//这一步其实就是两个有序数组的排序，nums[left:mid]和nums[mid+1:right]
	merge = func(nums []int, left, mid, right int) {
		tmp := make([]int, right-left+1)
		i, j, k := left, mid+1, 0
		for i <= mid && j <= right {
			if nums[i] < nums[j] {
				tmp[k] = nums[i]
				k++
				i++
			} else {
				tmp[k] = nums[j]
				k++
				j++
			}
		}
		for i <= mid {
			tmp[k] = nums[i]
			k++
			i++
		}
		for j <= right {
			tmp[k] = nums[j]
			k++
			j++
		}
		for i, v := range tmp {
			nums[left+i] = v
		}
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

//堆排序
func duipaixu(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	length := len(nums)
	var heapify func(nums []int, length, parent int)
	heapify = func(nums []int, length, parent int) {
		leftChild, rightChild := 2*parent+1, 2*parent+2
		//寻找父子节点中，最大的那个节点
		max := parent
		if leftChild < length && nums[leftChild] > nums[parent] {
			max = leftChild
		}
		if rightChild < length && nums[rightChild] > nums[parent] {
			max = rightChild
		}
		//最大的如果不是父节点，父节点下沉，子节点上升
		if max != parent {
			nums[max], nums[parent] = nums[parent], nums[max]
			//如果父节点比所有子节点都小，那刚才只替换了一个子节点，这里还需要在替换另一个子节点
			heapify(nums, length, max)
		}
	}
	//构建一个堆
	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, length, i)
	}
	for i := length - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}
	return nums
}
