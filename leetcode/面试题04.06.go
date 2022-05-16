package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历
//面试题 04.06. 后继者

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return inorderSuccessor_inOrder(root, p)          //中序遍历
	return inorderSuccessor_inOrder_iterator(root, p) //中序遍历-迭代
	return inorderSuccessor_bfs(root, p)              //利用BFS特性
}

//中序遍历
func inorderSuccessor_inOrder(root *TreeNode, p *TreeNode) *TreeNode {
	//找到了p
	isFind := false
	var result *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if result != nil {
			return
		}
		if result == nil && root.Left != nil {
			dfs(root.Left)
		}
		if result == nil && isFind {
			result = root
			return
		}
		if !isFind && root == p {
			isFind = true
		}
		if result == nil && root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return result
}

//中序遍历-迭代
func inorderSuccessor_inOrder_iterator(root *TreeNode, p *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	var pre, cur *TreeNode = nil, root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	return nil
}

//利用BFS的特性
func inorderSuccessor_bfs(root *TreeNode, p *TreeNode) *TreeNode {
	var result *TreeNode
	if p.Right != nil {
		result = p.Right
		for result.Left != nil {
			result = result.Left
		}
		return result
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			result = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return result
}
