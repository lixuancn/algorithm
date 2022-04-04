package main

//979. 在二叉树中分配硬币

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//后序遍历。题解https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/solution/by-yingying123-cif7/

var ans = 0

//一个节点的左子树需要移动left个，右子树需要移动left个，父节点有node.Val个，那么需要移动node.Val-1个。
//那么步骤需要多走abs(left) + abs(right)
//则一个树内部消化后（父左右已经自己挪完了），这时还多出来了left+right+node.Val-1个（可能正可能负）
func distributeCoins(root *TreeNode) int {
	ans = 0
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	ans += abs(left) + abs(right)
	return left + right + root.Val - 1
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}
