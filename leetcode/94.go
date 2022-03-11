package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversal_recursion(root) //递归
	//inorderTraversal_iteration(root) //迭代
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
	return nil
}
