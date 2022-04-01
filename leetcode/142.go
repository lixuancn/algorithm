package main

import "fmt"

//142. 环形链表 II

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//数学可以证明，1个指针从头结点到环入口点，会遇到慢指针从相遇点走到环入口点。
		//推倒过程可以看题解
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

func main() {
	//head := detectCycle(newList([]int{1, 2, 3, 4, 5}))
	head := detectCycle(newList([]int{1}))
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
