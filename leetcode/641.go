package main

import "fmt"

//641. 设计循环双端队列

type MyCircularDeque struct {
	cap   int
	head  int
	tail  int
	size  int
	queue []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{cap: k, queue: make([]int, k, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	//队列为空时，队头插入元素，队头队尾指针都要指向它
	if this.IsEmpty() {
		this.head = (this.head + 1) % this.cap
		this.tail = (this.tail + 1) % this.cap
		this.queue[this.head] = value
		this.size++
		return true
	}
	this.head = (this.head - 1 + this.cap) % this.cap
	this.queue[this.head] = value
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = (this.head + 1) % this.cap
		this.tail = (this.tail + 1) % this.cap
		this.queue[this.tail] = value
		this.size++
		return true
	}
	this.tail = (this.tail + 1) % this.cap
	this.queue[this.tail] = value
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = (this.head - 1) % this.cap
		this.tail = (this.tail - 1) % this.cap
		this.size--
		return true
	}
	this.head = (this.head + 1) % this.cap
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = (this.head - 1) % this.cap
		this.tail = (this.tail - 1) % this.cap
		this.size--
		return true
	}
	this.tail = (this.tail - 1 + this.cap) % this.cap
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.cap
}

func main() {
	//circularDeque := Constructor(3)           // 设置容量大小为3
	//fmt.Println(circularDeque.InsertLast(1))  // 返回 true
	//fmt.Println(circularDeque.InsertLast(2))  // 返回 true
	//fmt.Println(circularDeque.InsertFront(3)) // 返回 true
	//fmt.Println(circularDeque.InsertFront(4)) // 已经满了，返回 false
	//fmt.Println(circularDeque.GetRear())      // 返回 2
	//fmt.Println(circularDeque.IsFull())       // 返回 true
	//fmt.Println(circularDeque.DeleteLast())   // 返回 true
	//fmt.Println(circularDeque.InsertFront(4)) // 返回 true
	//fmt.Println(circularDeque.GetFront())     // 返回 4

	circularDeque := Constructor(4)           // 设置容量大小为4
	fmt.Println(circularDeque.InsertFront(9)) // 返回 true
	fmt.Println(circularDeque.DeleteLast())   // 返回 true
	fmt.Println(circularDeque.GetRear())      // 返回 -1
	fmt.Println(circularDeque.GetFront())     // 返回 -1
	fmt.Println(circularDeque.GetFront())     // 返回 -1
	fmt.Println(circularDeque.DeleteFront())  // 返回 false
	fmt.Println(circularDeque.InsertFront(6)) // 返回 true
	fmt.Println(circularDeque.InsertLast(5))  // 返回 true
	fmt.Println(circularDeque.InsertFront(9)) // 返回 true
	fmt.Println(circularDeque.GetFront())     // 返回 9
	fmt.Println(circularDeque.InsertFront(6)) // 返回 true
}
