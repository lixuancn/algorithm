package main

//92. 反转链表 II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween_1(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	prev := dummyNode
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummyNode.Next
}

//上面的方法是题解，很晦涩，就再写一个
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	prev := dummyNode
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	if prev.Next == nil {
		return dummyNode.Next
	}
	end := prev.Next
	for i := left; i < right; i++ {
		end = end.Next
	}
	next := end.Next
	end.Next = nil
	start := prev.Next
	prev.Next = reverse(start)
	start.Next = next
	return dummyNode.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
