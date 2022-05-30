package main

import "container/list"

//1022. 从根到叶的二进制数之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	//return sumRootToLeaf_dfs(root) //递归
	return sumRootToLeaf_iteration(root) //迭代
}

func sumRootToLeaf_dfs(root *TreeNode) int {
	result := 0
	var dfs func(root *TreeNode, digit int)
	dfs = func(root *TreeNode, digit int) {
		digit = digit<<1 | root.Val
		if root.Left != nil {
			dfs(root.Left, digit)
		}
		if root.Right != nil {
			dfs(root.Right, digit)
		}
		if root.Left == nil && root.Right == nil {
			result += digit
		}
	}
	dfs(root, 0)
	return result
}

func sumRootToLeaf_iteration(root *TreeNode) int {
	stack := list.New()
	result, digit := 0, 0
	var pre *TreeNode
	for root != nil || stack.Len() > 0 {
		for root != nil {
			digit = digit<<1 | root.Val
			stack.PushBack(root)
			root = root.Left
		}
		root = stack.Back().Value.(*TreeNode)
		if root.Right == nil || root.Right == pre {
			if root.Left == nil && root.Right == nil {
				result += digit
			}
			digit = digit >> 1
			stack.Remove(stack.Back())
			pre = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return result
}
