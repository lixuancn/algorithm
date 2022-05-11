package main

import (
	"strconv"
	"strings"
)

//449. 序列化和反序列化二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//一个数组，要还原成二叉树，需要前序+中序，或者后序+中序。但二叉搜索树来说，只要有前序或者后序就可以还原了

//后序遍历
func Constructor() CodecPostOrder {
	return CodecPostOrder{}
}

//前序遍历
// func Constructor() CodecPreOrder {
//     return CodecPreOrder{}
// }

//后序遍历
type CodecPostOrder struct {
}

// Serializes a tree to a single string.
func (this *CodecPostOrder) serialize(root *TreeNode) string {
	//后序遍历。数组的最后一个值就是根节点
	result := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
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
	if data == "" {
		return nil
	}
	list := strings.Split(data, " ")
	var dfs func(min, max int) *TreeNode
	dfs = func(min, max int) *TreeNode {
		if len(list) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(list[len(list)-1])
		//右边肯定比根大，左边肯定比根小
		if val < min || val > max {
			return nil
		}
		list = list[:len(list)-1]
		return &TreeNode{Val: val, Right: dfs(val, max), Left: dfs(min, val)}
	}
	return dfs(-1<<31, 1<<31)
}

//前序遍历
type CodecPreOrder struct {
}

// Serializes a tree to a single string.
func (this *CodecPreOrder) serialize(root *TreeNode) string {
	//前序遍历。数组的第一个值就是根节点
	result := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
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
	if data == "" {
		return nil
	}
	list := strings.Split(data, " ")
	val, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: val}
	leftList, rightList := []string{}, []string{}
	for i := 1; i < len(list); i++ {
		cur, _ := strconv.Atoi(list[i])
		if cur < val {
			leftList = append(leftList, list[i])
		} else {
			rightList = append(rightList, list[i])
		}
	}
	root.Left = this.deserialize(strings.Join(leftList, " "))
	root.Right = this.deserialize(strings.Join(rightList, " "))
	return root
}
