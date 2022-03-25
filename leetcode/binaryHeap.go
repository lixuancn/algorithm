package main

import "fmt"

type BinaryHeap struct {
	size int
	data []int
}

//大顶堆
func NewBinaryHeap(cap int) *BinaryHeap {
	data := make([]int, cap, cap)
	return &BinaryHeap{
		size: 0,
		data: data,
	}
}

func (this *BinaryHeap) isEmpty() bool {
	return this.size == 0
}

func (this *BinaryHeap) isFull() bool {
	return this.size == len(this.data)
}

func (this *BinaryHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (this *BinaryHeap) getChildIndex(parentIndex int) (int, int) {
	return 2*parentIndex + 1, 2*parentIndex + 2
}

func (this *BinaryHeap) insert(x int) {
	if this.isFull() {
		return
	}
	this.data[this.size] = x
	this.size++
	this.moveToUp(this.size - 1)
}

func (this *BinaryHeap) delete(x int) int {
	if this.isEmpty() {
		return -1
	}
	element := this.data[x]
	this.data[x] = this.data[this.size-1]
	this.size--
	this.moveToDown(x)
	return element
}

func (this *BinaryHeap) moveToUp(index int) {
	for index > 0 && this.data[index] > this.data[this.getParentIndex(index)] {
		this.data[index], this.data[this.getParentIndex(index)] = this.data[this.getParentIndex(index)], this.data[index]
		index = this.getParentIndex(index)
	}
}

func (this *BinaryHeap) moveToDown(index int) {
	for {
		maxChild := this.getMaxChild(index)
		if maxChild == -1 {
			break
		}
		if this.data[maxChild] <= this.data[index] {
			break
		}
		this.data[maxChild], this.data[index] = this.data[index], this.data[maxChild]
		index = maxChild
	}
}

func (this *BinaryHeap) getMaxChild(index int) int {
	leftChild, rightChild := this.getChildIndex(index)
	if leftChild >= this.size {
		return -1
	}
	if rightChild >= this.size {
		return leftChild
	}
	maxChild := leftChild
	if this.data[maxChild] < this.data[rightChild] {
		maxChild = rightChild
	}
	return maxChild
}

func main() {
	nums := []int{7, 3, 1, 2, 9, 4, 8, 6, 5}
	heap := NewBinaryHeap(len(nums))
	for _, num := range nums {
		heap.insert(num)
	}
	fmt.Println(heap.data)
	for range nums {
		fmt.Println(heap.data[0])
		heap.delete(0)
		fmt.Println(heap.data)
	}
}
