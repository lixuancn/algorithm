package main

import "fmt"

//链表相交

//用map，空间复杂度O(N)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cache := make(map[*ListNode]int)
	for headA != nil {
		cache[headA] = 1
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := cache[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

//用指针，空间复杂度O(1)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	lenA, lenB, step := 0, 0, 0
	long, slow := headA, headB
	for headA != nil {
		lenA++
		headA = headA.Next
	}
	for headB != nil {
		lenB++
		headB = headB.Next
	}
	if lenA > lenB {
		step = lenA - lenB
	} else {
		step = lenB - lenA
		long, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		long = long.Next
	}
	for long != slow {
		long = long.Next
		slow = slow.Next
	}
	return long
}

func main() {
	//head := getIntersectionNode(newList([]int{1, 2, 3, 4, 5}), 2)
	head := getIntersectionNode(newList([]int{1}), 1)
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
