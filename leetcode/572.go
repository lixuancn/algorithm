package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, subRoot) || check(root.Left, subRoot) || check(root.Right, subRoot)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil || q != nil {
		return false
	} else if p.Val != p.Val {
		return false
	}
	return check(p.Left, q.Left) && check(p.Right, q.Right)
}
