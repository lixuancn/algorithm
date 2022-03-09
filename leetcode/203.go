package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	node := dummyHead
	for node != nil && node.Next != nil {
		if val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return dummyHead.Next
}

func main() {
	//head := removeElements(newList([]int{1, 2, 6, 3, 4, 5, 6}), 6)
	head := removeElements(newList([]int{7, 7, 7, 7}), 7)
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
