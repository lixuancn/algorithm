package main

//2. 两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	prev := dummyNode
	var node *ListNode
	n1, n2 := l1, l2
	carry := 0 //上一位进过来的
	for n1 != nil || n2 != nil {
		v1, v2 := 0, 0
		if n1 != nil {
			v1 = n1.Val
			node = n1
			n1 = n1.Next
		}
		if n2 != nil {
			v2 = n2.Val
			node = n2
			n2 = n2.Next
		}
		val := v1 + v2 + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val = val % 10
		}
		node.Val = val
		prev.Next = node
	}
	if carry == 1 {
		prev.Next = &ListNode{Val: 1}
	}
	return dummyNode.Next
}
