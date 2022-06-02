package main

//450. 删除二叉搜索树中的节点

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		//当前节点就是要删除的节点
		//左边为空，直接连右边
		if root.Left == nil {
			return root.Right
		}
		//右边为空，直接连左边
		if root.Right == nil {
			return root.Left
		}
		//左右节点都有，找右边的最小值
		node := root.Right
		for node.Left != nil {
			node = node.Left
		}
		//删除当前节点，把当前节点的左树给它
		node.Left = root.Left
		//当前节点被替代
		root = root.Right
	}
	return root
}
