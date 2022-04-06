package main

import (
	"container/heap"
	"fmt"
)

//973. 最接近原点的 K 个点
//转换为求数组中的第k小值得问题

func kClosest(points [][]int, k int) [][]int {
	//quickSelect(points, k, 0, len(points)-1) //快速选择
	//return points[:k]
	//return heapifySort(points, k) //堆排序
	//方法三：直接用堆，维护k个节点的最大堆，前k个节点放进堆中，第k+1和头比较，如果小的话，就把头弹出。这样堆中就是最小的k个节点
	return heapify(points, k)
}

func getDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

//方法一 快速选择
func quickSelect(points [][]int, k, left, right int) {
	counter, pivot := left, right
	for i := left; i < right; i++ {
		if getDistance(points[i]) < getDistance(points[pivot]) {
			points[i], points[counter] = points[counter], points[i]
			counter++
		}
	}
	points[pivot], points[counter] = points[counter], points[pivot]
	if counter+1 == k {
		return
	} else if counter+1 > k {
		//去左边
		quickSelect(points, k, left, counter-1)
	} else {
		//去右边
		quickSelect(points, k, counter+1, right)
	}
}

//方法一 堆排序
func heapifySort(points [][]int, k int) [][]int {
	heapSize := len(points)
	buildMaxHeapify(points, heapSize)
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = points[0]
		points[0], points[heapSize-1] = points[heapSize-1], points[0]
		heapSize--
		maxHeapify(points, heapSize, 0)
	}
	return result
}

func buildMaxHeapify(point [][]int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(point, heapSize, i)
	}
}

func maxHeapify(points [][]int, heapSize int, i int) {
	left, right, least := 2*i+1, 2*i+2, i
	if left < heapSize && getDistance(points[left]) < getDistance(points[least]) {
		least = left
	}
	if right < heapSize && getDistance(points[right]) < getDistance(points[least]) {
		least = right
	}
	if least != i {
		points[i], points[least] = points[least], points[i]
		maxHeapify(points, heapSize, least)
	}
}

//方法三 语言自带的堆
type pair struct {
	dist  int
	point []int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func heapify(points [][]int, k int) [][]int {
	h := make(hp, k)
	for i := range points[:k] {
		h[i] = pair{dist: getDistance(points[i]), point: points[i]}
	}
	heap.Init(&h)
	for _, p := range points[k:] {
		dist := getDistance(p)
		if dist < h[0].dist {
			h[0] = pair{dist: dist, point: p}
			heap.Fix(&h, 0) //效率高于先pop再push
		}
	}
	result := make([][]int, k)
	for i, p := range h {
		result[i] = p.point
	}
	return result
}

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	//fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
