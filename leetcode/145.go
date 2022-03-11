package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return postorderTraversal_recursion(root) //递归
	//postorderTraversal_iteration(root) //迭代
}

func postorderTraversal_recursion(root *TreeNode) []int {
	var res = make([]int, 0)
	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node == nil {
			return
		}
		recursion(node.Left)
		recursion(node.Right)
		res = append(res, node.Val)
	}
	recursion(root)
	return res
}

func postorderTraversal_iteration(root *TreeNode) []int {
	return nil
}
