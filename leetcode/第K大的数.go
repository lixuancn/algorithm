package main

import "fmt"

//最大的第K个数

func main() {
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 9))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 8))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 7))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 6))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 5))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 4))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 3))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 2))
	fmt.Println(quickSelect([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 0, 8, 1))

	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 9))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 8))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 7))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 6))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 5))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 4))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 3))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 2))
	fmt.Println(heapifySort([]int{2, 1, 5, 9, 7, 8, 4, 6, 3}, 1))
}

func quickSelect(nums []int, start, end, k int) int {
	pivot := end
	counter := start
	for i := start; i < end; i++ {
		//正序
		if nums[i] <= nums[pivot] {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	if end-counter+1 == k {
		return nums[counter]
	}
	if end-counter+1 > k {
		//取右边
		return quickSelect(nums, counter+1, end, k)
	} else {
		//取左边
		return quickSelect(nums, start, counter-1, k-(end-counter+1))
	}
}

func heapifySort(nums []int, k int) int {
	//第一步，构建一个二叉树，数组就是可以想象成一个二叉树的层序遍历，这个就可以省略
	heapSize := len(nums)
	//第二步，调整二叉树，变成一个最大堆
	buildMaxHeap(nums, heapSize)
	//第三步，删除
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		//最后一个拿到头部来
		nums[0], nums[i] = nums[i], nums[0]
		//删除头，长度减一，就代表刚才把头换到尾巴后，就永远也不能被访问了，因为长度超出了，就相当于pop操作一样
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

//每个树（子树）的根节点看是否要向下走
func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	left, right, largest := 2*i+1, 2*i+2, i
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		maxHeapify(nums, largest, heapSize)
	}
}
