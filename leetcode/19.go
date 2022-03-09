package main

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	slowNode, quickNode := dummyNode, dummyNode
	for n >= 0 {
		quickNode = quickNode.Next
		n--
	}
	for quickNode != nil {
		quickNode = quickNode.Next
		slowNode = slowNode.Next
	}
	slowNode.Next = slowNode.Next.Next
	return dummyNode.Next
}

func main() {
	//head := removeNthFromEnd(newList([]int{1, 2, 3, 4, 5}), 2)
	head := removeNthFromEnd(newList([]int{1}), 1)
	fmt.Println(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(arr []int) *ListNode {
	var nextNode *ListNode = nil
	for i := len(arr) - 1; i >= 0; i-- {
		node := &ListNode{Val: arr[i], Next: nextNode}
		nextNode = node
	}
	return nextNode
}
