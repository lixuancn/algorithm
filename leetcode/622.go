package main

import "fmt"

//622. 设计循环队列
//只是说一个数组可以反复来用，并不是真正的一个环，因为真正的一个环不存在容量满了插入失败的情况，环是覆盖的

type MyCircularQueue struct {
	cap   int
	head  int
	size  int
	queue []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{cap: k, queue: make([]int, k, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[(this.head+this.size)%this.cap] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}
	this.head = (this.head + 1) % this.cap
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]

}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[(this.head+this.size-1)%this.cap]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}

func main() {
	circularQueue := Constructor(3)
	fmt.Println(circularQueue.EnQueue(1)) // 返回 true
	fmt.Println(circularQueue.EnQueue(2)) // 返回 true
	fmt.Println(circularQueue.EnQueue(3)) // 返回 true
	fmt.Println(circularQueue.EnQueue(4)) // 返回 false，队列已满
	fmt.Println(circularQueue.Rear())     // 返回 3
	fmt.Println(circularQueue.IsFull())   // 返回 true
	fmt.Println(circularQueue.DeQueue())  // 返回 true
	fmt.Println(circularQueue.EnQueue(4)) // 返回 true
	fmt.Println(circularQueue.Rear())     // 返回 4
}
