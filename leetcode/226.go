package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层序遍历（BFS广度优先）
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}

//前序遍历 递归
func invertTree_2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree_2(root.Left)
	invertTree_2(root.Right)
	return root
}

//前序遍历 迭代
func invertTree_3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		n := stack.Len()
		for i := 0; i < n; i++ {
			node := stack.Remove(stack.Back()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}
	}
	return root
}

//中序遍历 会交换两次，所以不合适。如果非要写，就两行都是对Left进行迭代（递归）

//后序遍历 递归
func invertTree_4(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree_2(root.Left)
	invertTree_2(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

//后序遍历 迭代
func invertTree_5(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		n := stack.Len()
		for i := 0; i < n; i++ {
			node := stack.Remove(stack.Back()).(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}
