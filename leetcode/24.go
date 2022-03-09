package main

import "fmt"

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n3 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = n1
		cur.Next.Next.Next = n3

		cur = cur.Next.Next
	}
	return dummyHead.Next
}

func main() {
	head := swapPairs(newList([]int{1, 2, 3, 4}))
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
