package main

//23. 合并K个升序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

//这是我自己想的办法。
//题解的方法一是转化为两个链表排序
//ans=&ListNode{}
//for _, list := range lists{
//	merge(ans, list)
//}
//题解的方法二是分治，类似归并的思路，可以练习一下
func mergeKLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}
	prev := dummyNode
	for {
		minVal, minIndex := 1<<31, -1
		notNilListCount := 0 //非空节点的个数，只剩一个的话就可以跳出while了
		for i, ln := range lists {
			if ln == nil {
				continue
			}
			notNilListCount++
			if ln.Val < minVal {
				minVal = ln.Val
				minIndex = i
			}
		}
		if minIndex == -1 {
			break
		}
		if notNilListCount == 1 {
			prev.Next = lists[minIndex]
			break
		}
		node := lists[minIndex]
		lists[minIndex] = node.Next
		node.Next = nil
		prev.Next = node
		prev = node
	}
	return dummyNode.Next
}

//归并
func mergeKLists_1(lists []*ListNode) *ListNode {
	return mergeSort(lists, 0, len(lists)-1)
}

func mergeSort(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	l1 := mergeSort(lists, left, mid)
	l2 := mergeSort(lists, mid+1, right)
	return merge(l1, l2)
}

func merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummyNode := &ListNode{}
	prev := dummyNode
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
	if list1 == nil {
		prev.Next = list2
	}
	if list2 == nil {
		prev.Next = list1
	}
	return dummyNode.Next
}
