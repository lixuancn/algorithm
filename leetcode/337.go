package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//337. 打家劫舍 III

func rob(root *TreeNode) int {
	return rob_dfs(root)
	return rob_dfs_cache(root)
	return rob_dp(root)
}

//暴力法 - 后序遍历 O(n^2)
//会超时，我们发现，有重复遍历的情况，根的左左子树和右右子树遍历的时候，会遍历的左左左子树和右右右子树。
//而根遍历左右子树的时候，也会遍历到根的左左子树和右右子树，还有根的左左左子树和右右右子树，这里可以用记忆化递归
func rob_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	//偷本节点
	val1 := root.Val
	if root.Left != nil {
		val1 += rob_dfs(root.Left.Left) + rob_dfs(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob_dfs(root.Right.Left) + rob_dfs(root.Right.Right)
	}
	//不偷本节点
	val2 := rob_dfs(root.Left) + rob_dfs(root.Right)
	return max(val1, val2)
}

//暴力法 - 后序遍历 - 记忆化递归 O(n)
//会超时，我们发现，有重复遍历的情况，根的左左子树和右右子树遍历的时候，会遍历的左左左子树和右右右子树。
//而根遍历左右子树的时候，也会遍历到根的左左子树和右右子树，还有根的左左左子树和右右右子树，这里可以用记忆化递归
func rob_dfs_cache(root *TreeNode) int {
	cache := map[*TreeNode]int{}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return root.Val
		}
		if _, ok := cache[root]; ok {
			return cache[root]
		}
		//偷本节点
		val1 := root.Val
		if root.Left != nil {
			val1 += rob_dfs(root.Left.Left) + rob_dfs(root.Left.Right)
		}
		if root.Right != nil {
			val1 += rob_dfs(root.Right.Left) + rob_dfs(root.Right.Right)
		}
		//不偷本节点
		val2 := dfs(root.Left) + dfs(root.Right)
		cache[root] = max(val1, val2)
		return cache[root]
	}
	return dfs(root)
}

//树形动态规划的入门
//dp[2] 0表示该节点不偷，1表示偷。
//dp[0]=max(left[0], left[1]) + max(right[0], right[1])
//dp[1]=val+left[0]+right[0]
//采用后序遍历的方式，每层直接返回即可，所以dp只需要2个值就可以了
func rob_dp(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		left0, left1 := dfs(root.Left)
		right0, right1 := dfs(root.Right)
		v0 := max(left0, left1) + max(right0, right1)
		v1 := root.Val + left0 + right0
		return v0, v1
	}
	v0, v1 := dfs(root)
	return max(v0, v1)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
