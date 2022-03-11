package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	//return preorderTraversal_recursion(root) //递归
	return preorderTraversal_iteration(root) //迭代
}

func preorderTraversal_recursion(root *TreeNode) []int {
	var res = make([]int, 0)
	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		recursion(node.Left)
		recursion(node.Right)
	}
	recursion(root)
	return res
}

func preorderTraversal_iteration(root *TreeNode) []int {
	var res = make([]int, 0)
	var stack = make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		res = append(res, node.Val)
	}
	return res
}
