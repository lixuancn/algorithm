package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	//return minDepth_iteration(root)
	return minDepth_recursion_1(root)
	return minDepth_recursion_2(root)
}

//迭代
func minDepth_iteration(root *TreeNode) int {
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
			if node.Left == nil && node.Right == nil {
				return depth
			}
		}
	}
	return depth
}

//后序
func minDepth_recursion_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return minDepth_recursion_1(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth_recursion_1(root.Left) + 1
	} else {
		return min(minDepth_recursion_1(root.Left), minDepth_recursion_1(root.Right)) + 1
	}
}

//前序
var result int

func minDepth_recursion_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result = 1 << 31
	recursion(root, 1)
	return result
}

func recursion(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		result = min(result, depth)
		return
	}
	if root.Left != nil {
		depth++
		recursion(root.Left, depth)
		depth--
	}
	if root.Right != nil {
		depth++
		recursion(root.Right, depth)
		depth--
	}
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
