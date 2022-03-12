package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

//递归
func postorder_recursion(root *Node) []int {
	var ret = make([]int, 0)
	var recursion func(root *Node)
	recursion = func(root *Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			recursion(child)
		}
		ret = append(ret, root.Val)
	}
	recursion(root)
	return ret
}

//迭代 - 除了前序的结果翻转，还可以如下操作
func postorder(root *Node) []int {
	var ret = make([]int, 0)
	if root == nil {
		return ret
	}
	stack := list.New()
	stack.PushBack(root)
	record := map[*Node]bool{}
	for stack.Len() > 0 {
		node := stack.Back().Value.(*Node)
		if len(node.Children) == 0 || record[node] {
			ret = append(ret, node.Val)
			stack.Remove(stack.Back())
			continue
		}
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
		record[node] = true
	}
	return ret
}
