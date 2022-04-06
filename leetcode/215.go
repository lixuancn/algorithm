package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	//return findKthLargest_quickSelect(nums, k) //快速选择
	return findKthLargest_heapifySort(nums, k) //堆排序
}

func findKthLargest_quickSelect(nums []int, k int) int {
	n := len(nums)
	var recursion func(nums []int, k int, left, right int) int
	recursion = func(nums []int, k int, left, right int) int {
		pivot, counter := right, left
		for i := left; i < right; i++ {
			if nums[i] < nums[pivot] {
				nums[i], nums[counter] = nums[counter], nums[i]
				counter++
			}
		}
		nums[counter], nums[pivot] = nums[pivot], nums[counter]
		if right-counter+1 == k {
			return nums[counter]
		} else if right-counter+1 < k {
			//在counter的左边去找
			return recursion(nums, k-(right-counter+1), left, counter-1)
		} else {
			//在counter的右边去找
			return recursion(nums, k, counter+1, right)
		}
	}
	return recursion(nums, k, 0, n-1)
}

//核心思路是堆排序后，删除顶部（最大值）操作k-1次后，堆顶就是第K大
//题解的图是很清晰的https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode-/
func findKthLargest_heapifySort(nums []int, k int) int {
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

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4
}
