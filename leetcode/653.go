package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return findTarget_dfs(root, k)
	return findTarget_iterator(root, k)
}

func findTarget_dfs(root *TreeNode, k int) bool {
	record := make(map[int]struct{})
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if _, ok := record[k-root.Val]; ok {
			return true
		}
		record[root.Val] = struct{}{}
		if root.Left != nil {
			if dfs(root.Left) {
				return true
			}
		}
		if root.Right != nil {
			if dfs(root.Right) {
				return true
			}
		}
		return false
	}
	return dfs(root)
}

func findTarget_iterator(root *TreeNode, k int) bool {
	record := make(map[int]struct{})
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Back()).(*TreeNode)
		if _, ok := record[k-node.Val]; ok {
			return true
		}
		record[node.Val] = struct{}{}
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return false
}
