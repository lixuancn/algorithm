package main

import "container/list"

//965. 单值二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return isUnivalTree_preorder(root)  //前序
	return isUnivalTree_midorder(root)  //中序
	return isUnivalTree_postorder(root) //后序
	return isUnivalTree_bfs(root)       //层序

}

func isUnivalTree_preorder(root *TreeNode) bool {
	if root == nil {
		return false
	}
	val := root.Val
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root.Val != val {
			return false
		}
		if root.Left != nil {
			if !dfs(root.Left) {
				return false
			}
		}
		if root.Right != nil {
			if !dfs(root.Right) {
				return false
			}
		}
		return true
	}
	return dfs(root)
}

func isUnivalTree_midorder(root *TreeNode) bool {
	if root == nil {
		return false
	}
	val := root.Val
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root.Left != nil {
			if !dfs(root.Left) {
				return false
			}
		}
		if root.Val != val {
			return false
		}
		if root.Right != nil {
			if !dfs(root.Right) {
				return false
			}
		}
		return true
	}
	return dfs(root)
}

func isUnivalTree_postorder(root *TreeNode) bool {
	if root == nil {
		return false
	}
	val := root.Val
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root.Left != nil {
			if !dfs(root.Left) {
				return false
			}
		}
		if root.Right != nil {
			if !dfs(root.Right) {
				return false
			}
		}
		if root.Val != val {
			return false
		}
		return true
	}
	return dfs(root)
}

func isUnivalTree_bfs(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	queue.PushBack(root)
	val := root.Val
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		if node.Val != val {
			return false
		}
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return true
}
