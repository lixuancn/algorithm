package main

import (
	"fmt"
)

//2个栈来模拟队列。先入A栈，然后A栈出再入到B栈，这样就是队列的顺序了。

type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			val := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, val)
		}
	}
	if len(this.s2) == 0 {
		return 0
	}
	res := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			val := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, val)
		}
	}
	if len(this.s2) == 0 {
		return 0
	}
	return this.s2[len(this.s2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[len(s)-1])
	s = s[:len(s)-1]
	fmt.Println(s)
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	//fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}
