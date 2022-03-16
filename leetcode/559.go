package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	return maxDepth_iteration(root)
	return maxDepth_recursion_1(root)
	return maxDepth_recursion_2(root)
}

//迭代
func maxDepth_iteration(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		depth++
		for i := 0; i < n; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
			for _, node := range cur.Children {
				if node != nil {
					queue.PushBack(node)
				}
			}
		}
	}
	return depth
}

//后序
func maxDepth_recursion_1(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for _, node := range root.Children {
		depth = max(depth, maxDepth_recursion_1(node))
	}
	return depth + 1
}

//前序
var result int

func maxDepth_recursion_2(root *Node) int {
	if root == nil {
		return 0
	}
	result = 0
	recursion(root, 1)
	return result
}

func recursion(root *Node, depth int) {
	result = max(result, depth)
	if root == nil {
		return
	}
	for _, child := range root.Children {
		if child != nil {
			depth++
			recursion(child, depth)
			depth--
		}
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
