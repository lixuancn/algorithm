package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTree_recursion(p, q)
	return isSameTree_iteration(p, q)
}

func isSameTree_recursion(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree_recursion(p.Left, q.Left) && isSameTree_recursion(p.Right, q.Right)
}

func isSameTree_iteration(p *TreeNode, q *TreeNode) bool {
	queue := list.New()
	queue.PushBack(p)
	queue.PushBack(q)
	for queue.Len() > 0 {
		p = queue.Remove(queue.Front()).(*TreeNode)
		q = queue.Remove(queue.Front()).(*TreeNode)
		if p == nil && q == nil {
			continue
		} else if p == nil || q == nil {
			return false
		} else if p.Val != q.Val {
			return false
		}
		queue.PushBack(p.Left)
		queue.PushBack(q.Left)
		queue.PushBack(p.Right)
		queue.PushBack(q.Right)
	}
	return true
}
