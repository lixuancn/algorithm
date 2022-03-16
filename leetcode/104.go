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
	//return maxDepth_iteration(root)
	return maxDepth_recursion_1(root)
	return maxDepth_recursion_2(root)
}

//迭代 - 层序遍历
func maxDepth_iteration(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		depth++
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
	d1 := maxDepth_recursion_1(root.Left)
	d2 := maxDepth_recursion_1(root.Right)
	return max(d1, d2) + 1
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
	if result < depth {
		result = depth
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil {
		depth++
		front_recursion(root.Left, depth)
		depth--
	}
	if root.Right != nil {
		depth++
		front_recursion(root.Right, depth)
		depth--
	}
}
