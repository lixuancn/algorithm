package main

//剑指 Offer 54. 二叉搜索树的第k大节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res, k int

func kthLargest(root *TreeNode, k int) int {
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	k--
	if k == 0 {
		res = root.Val
		return
	}
	dfs(root.Left)
}
