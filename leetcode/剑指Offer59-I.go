package main

//维护一个单调递增队列

type IncrQueue struct {
	queue []int
}

func NewIncrQueue() *IncrQueue {
	return &IncrQueue{}
}

func (q *IncrQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *IncrQueue) Back() int {
	if q.Empty() {
		panic("递增队列为空")
	}
	return q.queue[len(q.queue)-1]
}

func (q *IncrQueue) Front() int {
	if q.Empty() {
		panic("递增队列为空")
	}
	return q.queue[0]
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

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
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
