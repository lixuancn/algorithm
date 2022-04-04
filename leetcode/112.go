package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//112. 路径总和

func hasPathSum(root *TreeNode, targetSum int) bool {
	//return hasPathSum_iterator(root, targetSum)
	return hasPathSum_dfs(root, targetSum)
}

//迭代
func hasPathSum_iterator(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodeQueue := list.New()
	valQueue := list.New()
	nodeQueue.PushBack(root)
	valQueue.PushBack(root.Val)
	for nodeQueue.Len() > 0 {
		node := nodeQueue.Remove(nodeQueue.Front()).(*TreeNode)
		val := valQueue.Remove(valQueue.Front()).(int)
		if node.Left == nil && node.Right == nil {
			if val == targetSum {
				return true
			}
		}
		if node.Left != nil {
			nodeQueue.PushBack(node.Left)
			valQueue.PushBack(val + node.Left.Val)
		}
		if node.Right != nil {
			nodeQueue.PushBack(node.Right)
			valQueue.PushBack(val + node.Right.Val)
		}
	}
	return false
}

//递归，深搜
func hasPathSum_dfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum_dfs(root.Left, targetSum-root.Val) || hasPathSum_dfs(root.Right, targetSum-root.Val)
}
