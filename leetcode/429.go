package main

import (
	"container/list"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		resLine := make([]int, 0)
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for _, child := range node.Children {
				queue.PushBack(child)
			}
			resLine = append(resLine, node.Val)
		}
		res = append(res, resLine)
	}
	return res
}
