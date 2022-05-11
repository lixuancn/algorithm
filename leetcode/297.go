package main

import (
	"strconv"
	"strings"
)

//297. 二叉树的序列化与反序列化

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func Constructor() CodecPreOrder {
	return CodecPreOrder{}
}

//后序遍历
//func Constructor() CodecPostOrder {
//	return CodecPostOrder{}
//}

//前序遍历
type CodecPreOrder struct {
}

// Serializes a tree to a single string.
func (this *CodecPreOrder) serialize(root *TreeNode) string {
	//前序遍历
	result := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			result = append(result, "null")
			return
		}
		result = append(result, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(result, " ")
}

// Deserializes your encoded data to tree.
func (this *CodecPreOrder) deserialize(data string) *TreeNode {
	list := strings.Split(data, " ")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if list[0] == "null" {
			list = list[1:]
			return nil
		}
		val, _ := strconv.Atoi(list[0])
		list = list[1:]
		return &TreeNode{Val: val, Left: dfs(), Right: dfs()}
	}
	return dfs()
}

//后序遍历
type CodecPostOrder struct {
}

// Serializes a tree to a single string.
func (this *CodecPostOrder) serialize(root *TreeNode) string {
	//前序遍历
	result := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			result = append(result, "null")
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		result = append(result, strconv.Itoa(root.Val))
	}
	dfs(root)
	return strings.Join(result, " ")
}

// Deserializes your encoded data to tree.
func (this *CodecPostOrder) deserialize(data string) *TreeNode {
	list := strings.Split(data, " ")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if list[len(list)-1] == "null" {
			list = list[1:]
			return nil
		}
		val, _ := strconv.Atoi(list[len(list)-1])
		list = list[:len(list)-1]
		return &TreeNode{Val: val, Right: dfs(), Left: dfs()}
	}
	return dfs()
}
