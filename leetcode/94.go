package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	//return inorderTraversal_recursion(root) //递归
	return inorderTraversal_iteration(root) //迭代
}

func inorderTraversal_recursion(root *TreeNode) []int {
	var res = make([]int, 0)
	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node == nil {
			return
		}
		recursion(node.Left)
		res = append(res, node.Val)
		recursion(node.Right)
	}
	recursion(root)
	return res
}

func inorderTraversal_iteration(root *TreeNode) []int {
	var res = make([]int, 0)
	if root == nil {
		return res
	}
	var stack = list.New()
	cur := root
	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
