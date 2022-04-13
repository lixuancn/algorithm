package main

import "fmt"

//75. 颜色分类

func sortColors(nums []int) {
	//sortColors_quickSelect(nums) //快排
	//sortColors_mergeSort(nums) //归并
	//sortColors_heaipy(nums) //归并排序
	//sortColors_countSort(nums) //计数排序
	//sortColors_pointer(nums) //计数排序优化版，单指针
	//sortColors_doublePointer(nums) //单指针优化版，双指针
	sortColors_doublePointer_1(nums) //双指针，另一种方法
}

//快排
func sortColors_quickSelect(nums []int) {
	var quickSelect func(nums []int, left, right int)
	quickSelect = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		counter, pivot := left, right
		for i := left; i < right; i++ {
			if nums[i] <= nums[pivot] {
				nums[counter], nums[i] = nums[i], nums[counter]
				counter++
			}
		}
		nums[counter], nums[pivot] = nums[pivot], nums[counter]
		quickSelect(nums, left, counter-1)
		quickSelect(nums, counter+1, right)
	}
	quickSelect(nums, 0, len(nums)-1)
}

//归并
func sortColors_mergeSort(nums []int) {
	var mergeSort func(nums []int, left, right int)
	var merge func(nums []int, left, mid, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		mid := (left + right) >> 1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
	merge = func(nums []int, left, mid, right int) {
		tmp := make([]int, right-left+1)
		i, j, k := left, mid+1, 0
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				tmp[k] = nums[i]
				i++
			} else {
				tmp[k] = nums[j]
				j++
			}
			k++
		}
		for ; i <= mid; i, k = i+1, k+1 {
			tmp[k] = nums[i]
		}
		for ; j <= right; j, k = j+1, k+1 {
			tmp[k] = nums[j]
		}
		for i := left; i <= right; i++ {
			nums[i] = tmp[i-left]
		}
		//for i, v := range tmp {
		//	nums[left+i] = v
		//}
	}
	mergeSort(nums, 0, len(nums)-1)
}

//归并
func sortColors_heaipy(nums []int) {
	var heapify func(nums []int, n int, i int)
	heapify = func(nums []int, n int, i int) {
		left, right, largest := 2*i+1, 2*i+2, i
		if left < n && nums[left] > nums[largest] {
			largest = left
		}
		if right < n && nums[right] > nums[largest] {
			largest = right
		}
		if largest != i {
			nums[largest], nums[i] = nums[i], nums[largest]
			heapify(nums, n, largest)
		}
	}
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}
}

//只包含0、1、2所以用计数排序很合适
func sortColors_countSort(nums []int) {
	counts := [3]int{}
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}
	i := 0
	for num, count := range counts {
		for k := 0; k < count; k++ {
			nums[i] = num
			i++
		}
	}
}

//计数排序会用到空间，优化为不用额外空间，用单指针，把所有0放到开头，再把所有1放到开头
func sortColors_pointer(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
}

//单指针需要循环2次，想减少到一次，2个指针p0,p1。
//如果是1，p1移动
//如果是0，就放到p0,此时p0p1同时移动。如果没有遇到1，那么p0p1一直在一起
//如果是0，如果p0<p1，那么现在要换的位置就是已经换好的1了，所以会把这个1换走，因此需要再做一次替换，把这个1再换到p1上去
func sortColors_doublePointer(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		} else if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		}
	}
}

//类似上面的双指针，我们用p0和p2指向两头
//0就放到p0, 2就放到p2, 再看p2换回来的可能是2/1/0，需要再次交换才行
func sortColors_doublePointer_1(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		} else if nums[i] == 2 {
			for nums[i] == 2 && i <= p2 {
				nums[i], nums[p2] = nums[p2], nums[i]
				p2--
			}
			if nums[i] == 0 && i <= p2 {
				nums[i], nums[p0] = nums[p0], nums[i]
				p0++
			}
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
