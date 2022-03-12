package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

//递归
func preorder_recursion(root *Node) []int {
	var ret = make([]int, 0)
	var recursion func(root *Node)
	recursion = func(root *Node) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		for _, child := range root.Children {
			recursion(child)
		}
	}
	recursion(root)
	return ret
}

//迭代
func preorder(root *Node) []int {
	var ret = make([]int, 0)
	if root == nil {
		return ret
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*Node)
		ret = append(ret, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}
	return ret
}
