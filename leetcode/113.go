package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	result = make([][]int, 0)
	pathSum_iterator(root, targetSum)
	//pathSum_dfs(root, targetSum)
	return result
}

//迭代
func pathSum_iterator(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return result
	}
	nodeQueue := list.New()
	nodeQueue.PushBack(root)
	valQueue := list.New()
	valQueue.PushBack(root.Val)
	ret := []int{root.Val}
	pathQueue := list.New()
	pathQueue.PushBack(ret)
	for nodeQueue.Len() > 0 {
		node := nodeQueue.Remove(nodeQueue.Front()).(*TreeNode)
		val := valQueue.Remove(valQueue.Front()).(int)
		ret = pathQueue.Remove(pathQueue.Front()).([]int)
		if node.Left == nil && node.Right == nil {
			if val == targetSum {
				result = append(result, append([]int(nil), ret...))
			}
		}
		if node.Left != nil {
			nodeQueue.PushBack(node.Left)
			valQueue.PushBack(val + node.Left.Val)
			pathQueue.PushBack(append(append([]int(nil), ret...), node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue.PushBack(node.Right)
			valQueue.PushBack(val + node.Right.Val)
			pathQueue.PushBack(append(append([]int(nil), ret...), node.Right.Val))
		}
	}
	return result
}

//递归，深搜
func pathSum_dfs(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	pathSum_dfs_handler(root, targetSum-root.Val, []int{root.Val})
}

func pathSum_dfs_handler(root *TreeNode, targetSum int, ret []int) {
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			tmp := make([]int, len(ret))
			copy(tmp, ret)
			result = append(result, tmp)
		}
		return
	}
	if root.Left != nil {
		ret = append(ret, root.Left.Val)
		pathSum_dfs_handler(root.Left, targetSum-root.Left.Val, ret)
		ret = ret[:len(ret)-1]
	}
	if root.Right != nil {
		ret = append(ret, root.Right.Val)
		pathSum_dfs_handler(root.Right, targetSum-root.Right.Val, ret)
		ret = ret[:len(ret)-1]
	}
}
