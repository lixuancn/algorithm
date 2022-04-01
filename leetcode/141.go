package main

//141. 环形链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, quick := head, head.Next
	for slow != nil && quick != nil {
		if slow == quick {
			return true
		}
		slow = slow.Next
		quick = quick.Next
		if quick == nil {
			return false
		}
		quick = quick.Next
	}
	return false
}
