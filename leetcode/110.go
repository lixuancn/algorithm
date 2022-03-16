package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//求深度 从上往下，适合前序
//求高度 从下往上，适合后序

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := dfs(root)
	if result == -1 {
		return false
	}
	return true
}

//高度是从下往上走，所以用后序遍历
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height1 := dfs(root.Left)
	if height1 == -1 {
		return -1
	}
	height2 := dfs(root.Right)
	if height2 == -1 {
		return -1
	}
	if height2-height1 > 1 || height1-height2 > 1 {
		return -1
	}
	return max(height1, height2) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
