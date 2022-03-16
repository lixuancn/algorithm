package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	//广度优先
	//return isSymmetric_bfs(root)
	//递归
	//return isSymmetric_recursion(root)
	//迭代
	return isSymmetric_iteration(root)
}

//BFS 层序遍历，判断每一行是否对称。
//自己想的，没有题解
func isSymmetric_bfs(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		line := make([]*TreeNode, n)
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			line[i] = node
			if node == nil {
				continue
			}
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		}

		n = len(line)
		if n > 1 && n&1 == 1 { //n&1==1 => n%2=1
			return false
		}
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			if line[i] == nil && line[j] == nil {
				continue
			} else if line[i] != nil && line[j] != nil {
				if line[i].Val != line[j].Val {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

//递归
//左子树的左节点和右子树的右节点比较，//左子树的右节点和右子树的左节点比较
func isSymmetric_recursion(root *TreeNode) bool {
	return recursion(root.Left, root.Right)
}

func recursion(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	//左子树的左节点和右子树的右节点比较
	r1 := recursion(left.Left, right.Right)
	//左子树的右节点和右子树的左节点比较
	r2 := recursion(left.Right, right.Left)
	//左右子树都对称，则该节点对称
	return r1 && r2
}

//迭代
func isSymmetric_iteration(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		left, right := queue.Remove(queue.Front()).(*TreeNode), queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		queue.PushBack(left.Left)
		queue.PushBack(right.Right)

		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return false
}
