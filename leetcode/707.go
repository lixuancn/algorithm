package main

import "fmt"

//未通过，不想调了。

type MyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

func Constructor() MyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	head.Pre = nil
	tail.Next = nil
	tail.Pre = head
	return MyLinkedList{head: head, tail: tail}
}

func (this *MyLinkedList) Get(index int) int {
	val := -1
	node := this.head.Next
	i := 0
	for node != nil && node.Next != nil {
		if i == index {
			val = node.Val
			break
		}
		node = node.Next
		i++
	}
	return val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{}
	node.Val = val
	node.Pre = this.head
	node.Next = this.head.Next
	this.head.Next.Pre = node
	this.head.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val, Pre: this.tail.Pre, Next: this.tail}
	this.tail.Pre.Next = node
	this.tail.Pre = node

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := this.head.Next
	i := 0
	for node != nil && node.Next != nil {
		if i == index {
			break
		}
		node = node.Next
		i++
	}
	if i != index {
		return
	}
	n := &Node{
		Next: node,
		Pre:  node.Pre,
		Val:  val,
	}
	node.Pre.Next = n
	node.Pre = n
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this.head.Next
	i := 0
	for node != nil && node.Next != nil {
		if i == index {
			break
		}
		node = node.Next
		i++
	}
	if i != index {
		return
	}
	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next
}

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	fmt.Println(linkedList.Get(0)) //返回1
	linkedList.AddAtTail(3)
	fmt.Println(linkedList.Get(1)) //返回3
	linkedList.AddAtIndex(1, 2)    //链表变为1-> 2-> 3
	fmt.Println(linkedList.Get(1)) //返回2
	linkedList.DeleteAtIndex(1)    //现在链表是1-> 3
	fmt.Println(linkedList.Get(1)) //返回3

	head := linkedList.head.Next
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
