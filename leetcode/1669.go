package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	start, end := list1, list1
	for i := 0; i < a-1; i++ {
		start = start.Next
	}
	for i := 0; i < b+1; i++ {
		end = end.Next
	}
	start.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = end

	return list1
}
