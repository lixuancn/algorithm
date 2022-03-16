package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return countNodes_wanquanerchashu(root) //完全二叉树特性来做
	// return countNodes_iteration(root) //迭代
	// return countNodes_recursion_1(root) //前序遍历
	// return countNodes_recursion_2(root) //中序遍历
	// return countNodes_recursion_3(root) //后序遍历
}

//利用完全二叉树的特性
func countNodes_wanquanerchashu(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right := root.Right
	leftHeight := 0
	rightHeight := 0
	for left != nil {
		left = left.Left
		leftHeight++
	}
	for right != nil {
		right = right.Right
		rightHeight++
	}
	if leftHeight == rightHeight {
		return (2 << leftHeight) - 1 //2^n-1
	}
	return countNodes_wanquanerchashu(root.Left) + countNodes_wanquanerchashu(root.Right) + 1
}

func countNodes_iteration(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	count := 0
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			count++
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return count
}

var count = 0

//前序
func countNodes_recursion_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count = 0
	recursion_1(root)
	return count
}

func recursion_1(root *TreeNode) {
	count++
	if root.Left != nil {
		recursion_1(root.Left)
	}
	if root.Right != nil {
		recursion_1(root.Right)
	}
}

//中序
func countNodes_recursion_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count = 0
	recursion_2(root)
	return count
}
func recursion_2(root *TreeNode) {
	if root.Left != nil {
		recursion_2(root.Left)
	}
	count++
	if root.Right != nil {
		recursion_2(root.Right)
	}
}

//后序
func countNodes_recursion_3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count = 0
	recursion_3(root)
	return count
}
func recursion_3(root *TreeNode) {
	if root.Left != nil {
		recursion_3(root.Left)
	}
	if root.Right != nil {
		recursion_3(root.Right)
	}
	count++
}
