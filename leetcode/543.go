package main

//543. 二叉树的直径
//求一棵二叉树中，距离最远的2个节点之间的距离

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func diameterOfBinaryTree(root *TreeNode) int {
	ret = 1
	dfs(root)
	return ret - 1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	ret = max(ret, l+r+1)
	return max(l, r) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
