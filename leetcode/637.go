package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var sum float64
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			sum += float64(node.Val)
		}
		res = append(res, sum/float64(n))
	}
	return res
}
