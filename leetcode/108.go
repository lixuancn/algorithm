package main

//108. 将有序数组转换为二叉搜索树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return recursion(nums, 0, len(nums)-1)
}

func recursion(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)>>1
	root := &TreeNode{Val: nums[mid]}
	root.Left = recursion(nums, left, mid-1)
	root.Right = recursion(nums, mid+1, right)
}
