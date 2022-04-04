package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return maxDepth_iteration(root)
	return maxDepth_recursion_1(root)
	return maxDepth_recursion_2(root)
}

//迭代 - 层序遍历
func maxDepth_iteration(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		depth++
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return depth
}

//后序遍历
func maxDepth_recursion_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_recursion_1(root.Left), maxDepth_recursion_1(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//前序遍历
var result int

func maxDepth_recursion_2(root *TreeNode) int {
	result = 0
	if root == nil {
		return result
	}
	front_recursion(root, 1)
	return result
}

func front_recursion(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if result < depth {
		result = depth
	}
	if root.Left != nil {
		front_recursion(root.Left, depth+1)
	}
	if root.Right != nil {
		front_recursion(root.Right, depth+1)
	}
}
