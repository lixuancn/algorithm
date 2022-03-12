package main

import (
	"container/list"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		var preNode *Node
		var node *Node
		for i := 0; i < n; i++ {
			if i == 0 {
				preNode = queue.Remove(queue.Front()).(*Node)
				node = preNode
			} else {
				node = queue.Remove(queue.Front()).(*Node)
				preNode.Next = node
				preNode = node
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		preNode.Next = nil
	}
	return root
}
