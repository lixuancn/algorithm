package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//k大于n时，k%n后，才是要需要移动的次数。因为移动n下就恢复了原样
//方法一：先把链表连成环，然后从尾巴处走n-k下，就是新的结尾，把这里把环断开
//方法二：快慢指针，快指针到头后，慢指针就是新结尾。
func rotateRight(head *ListNode, k int) *ListNode {
	//return rotateRight_ring(head, k)
	return rotateRight_doublePointer(head, k) //快慢指针
}

func rotateRight_ring(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	if k == n {
		return head
	}
	add := n - k%n
	tail.Next = head
	for add > 0 {
		add--
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}

func rotateRight_doublePointer(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	if k == n {
		return head
	}
	k = k % n
	//快指针先走k步
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	//快慢一起走
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}
