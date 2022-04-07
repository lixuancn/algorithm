package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//1,2,3,4,5,6,7
	//2,1,4,3,6,5,7
	dummyNode := &ListNode{}

	prev := dummyNode
	tow(prev, root)

	fmt.Println(dummyNode.Next)
}

func tow(prev, start *ListNode) {
	//没有了，或者还剩一个
	if start == nil || start.Next == nil {
		return
	}
	n1 := start.Next
	n2 := start.Next.Next

	n1.Next = start
	start.Next = n2
	prev.Next = n1

	prev = start
	tow(start, n2)
}

func merge(list1, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	prev := dummyNode
	//1,2,3,4    1,2,3
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	}
	if list2 != nil {
		prev.Next = list2
	}
	return dummyNode.Next
}
