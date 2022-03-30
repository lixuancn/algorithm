package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	prev, end := dummyNode, dummyNode

	for {
		//检测这一段是否是k个，少于k就直接结束
		i := 0
		for i = 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		//这一段的结尾end，和一段的开头，断开
		next := end.Next
		end.Next = nil
		start := prev.Next
		//翻转这一段，返回翻转后的头
		prev.Next = reverse(start)
		//和下一段链接
		start.Next = next
		//指针移动
		prev = start
		end = start
	}
	return dummyNode.Next
}

func reverse(root *ListNode) *ListNode {
	var prev *ListNode
	curr := root
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	root := &ListNode{Val: nums[0]}
	node := root
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{Val: nums[i]}
		node = node.Next
	}

	n := root
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}

	fmt.Println("------------")

	head := reverseKGroup(root, 3)

	n = head
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
