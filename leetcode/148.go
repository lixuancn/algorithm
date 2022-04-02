package main

import (
	"fmt"
	"strconv"
)

//148. 排序链表
//链表的归并排序

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sortList_mergeSort(head, nil) //归并
	return sortList_force(head)          //暴力法-插入排序
}

func sortList_mergeSort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return sortList_merge(sortList_mergeSort(head, slow), sortList_mergeSort(slow, tail))
}

func sortList_merge(list1, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}
	if list1 != nil {
		tmp.Next = list1
	}
	if list2 != nil {
		tmp.Next = list2
	}
	return dummyHead.Next
}

func sortList_force(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := head.Next
	root := head
	root.Next = nil
	for node != nil {
		//当前节点断开
		cur := node
		node = node.Next
		cur.Next = nil
		//把当前节点插入到新列表中
		root = sortInsert(root, cur)
	}
	return root
}
func printListNode(head *ListNode) {
	node := head
	str := ""
	for node != nil {
		str += strconv.Itoa(node.Val) + "->"
		node = node.Next
	}
	fmt.Println(str)
}
func sortInsert(head, target *ListNode) *ListNode {
	if target.Val < head.Val {
		target.Next = head
		return target
	}
	var prev *ListNode
	node := head
	for node != nil {
		if target.Val < node.Val {
			prev.Next = target
			target.Next = node
			return head
		}
		prev = node
		node = node.Next
	}
	//插不进去就跟到最后面
	prev.Next = target
	return head
}
