package main

import "fmt"

//155. 最小栈

//维护两个栈，一个是数据栈，一个是最小值的栈。同步插入，同步删除
//比如一个栈是abc，最小值是c的话，最小栈就是ccc。只要c还在数据栈里，c就是最小值
type MinStack struct {
	dataStack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{dataStack: make([]int, 0), minStack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.dataStack = append(this.dataStack, val)
	minVal := val
	if len(this.minStack) > 0 && minVal > this.minStack[len(this.minStack)-1] {
		minVal = this.minStack[len(this.minStack)-1]
	}
	this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	if len(this.dataStack) <= 0 {
		return
	}
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.dataStack) <= 0 {
		return 0
	}
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) <= 0 {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) //返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top())    // 返回 0.
	fmt.Println(minStack.GetMin()) // 返回 -2.
}
