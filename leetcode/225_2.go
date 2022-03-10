package main

//一个队列的方法

type MyStack struct {
	q1 []int
}

func Constructor() MyStack {
	return MyStack{
		q1: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	n := len(this.q1)
	this.q1 = append(this.q1, x)
	for i := 0; i < n; i++ {
		this.q1 = append(this.q1, this.q1[0])
		this.q1 = this.q1[1:]
	}
}

func (this *MyStack) Pop() int {
	v := this.q1[0]
	this.q1 = this.q1[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
