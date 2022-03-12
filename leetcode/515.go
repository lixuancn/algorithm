package main

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		max := math.MinInt32
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if max < node.Val {
				max = node.Val
			}
		}
		res = append(res, max)
	}
	return res
}
