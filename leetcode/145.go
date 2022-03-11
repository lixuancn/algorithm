package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	//return postorderTraversal_recursion(root) //递归
	return postorderTraversal_iteration(root) //迭代
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
	//前序是中左右，后序是左右中，那么把前序改成中右左然后翻转即可
	var res = make([]int, 0)
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}
