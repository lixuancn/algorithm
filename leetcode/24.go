package main

import "fmt"

//24. 两两交换链表中的节点

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

func swapPairs_practice(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	prev := dummyHead
	if prev == nil {
		return head
	}
	var n1, n2 *ListNode
	for prev.Next != nil && prev.Next.Next != nil {
		n1 = prev.Next
		n2 = prev.Next.Next
		n1.Next = n2.Next
		prev.Next = n2
		n2.Next = n1
		prev = n1
	}
	return dummyHead.Next
}
