package main

import "fmt"

//933. 最近的请求次数

//用队列来模拟
//之前我写的第一版维护一个3000长度的数组，每次写在t%3000的位置上。 然后每次遍历3000数组，看时间在[t-3000,t]之间的就++，最后返回

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	for this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))    //1
	fmt.Println(obj.Ping(100))  //2
	fmt.Println(obj.Ping(3001)) //3
	fmt.Println(obj.Ping(3002)) //3
	fmt.Println(obj.Ping(9003)) //1
}
