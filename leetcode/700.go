package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	//递归
	//return dfs(root, val)
	//迭代
	return iteration(root, val)
}

func dfs(node *TreeNode, val int) *TreeNode {
	if node == nil || node.Val == val {
		return node
	}
	if node.Val > val {
		return dfs(node.Left, val)
	} else {
		return dfs(node.Right, val)
	}
}

func iteration(node *TreeNode, val int) *TreeNode {
	for node != nil {
		if node.Val == val {
			return node
		} else if node.Val > val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
