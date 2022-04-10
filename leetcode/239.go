package main

import "fmt"

//239. 滑动窗口最大值

//封装一个单调队列
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (q *MyQueue) Front() int {
	return q.queue[0]
}

func (q *MyQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *MyQueue) Push(v int) {
	for !q.Empty() && q.Back() < v {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, v)
}

func (q *MyQueue) Pop(v int) {
	if !q.Empty() && q.Front() == v {
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0)
	q := NewMyQueue()
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	ret = append(ret, q.Front())
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		ret = append(ret, q.Front())
	}
	return ret
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow_practice([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

//封装一个单调队列
type IncrQueue struct {
	queue []int
}

func NewIncrQueue() *IncrQueue {
	return &IncrQueue{}
}

func (q *IncrQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *IncrQueue) Front() int {
	return q.queue[0]
}

func (q *IncrQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *IncrQueue) Push(v int) {
	for !q.Empty() && q.Back() < v {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, v)
}

func (q *IncrQueue) Pop(v int) {
	if !q.Empty() && q.Front() == v {
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindow_practice(nums []int, k int) []int {
	queue := NewIncrQueue()
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	result := make([]int, len(nums)-k+1)
	result[0] = queue.Front()
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		result[i-k+1] = queue.Front()
	}
	return result
}
