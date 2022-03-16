package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var rowList [][]string
var row []string

func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePaths_recursion(root)
	return binaryTreePaths_iteration(root)
}

func binaryTreePaths_recursion(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	rowList = make([][]string, 0)
	row = make([]string, 0)
	dfs(root)
	result := make([]string, len(rowList))
	for i, row := range rowList {
		result[i] = strings.Join(row, "->")
	}
	fmt.Println(result)
	return result
}

func dfs(root *TreeNode) {
	row = append(row, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		tmp := make([]string, len(row))
		copy(tmp, row)
		rowList = append(rowList, tmp)
		return
	}
	if root.Left != nil {
		dfs(root.Left)
		row = row[:len(row)-1]
	}
	if root.Right != nil {
		dfs(root.Right)
		row = row[:len(row)-1]
	}
}

func binaryTreePaths_iteration(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	row = make([]string, 0)
	result := make([]string, 0)
	stack := list.New()
	stack.PushBack(root)
	row = append(row, strconv.Itoa(root.Val))
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		path := row[len(row)-1]
		row = row[:len(row)-1]
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
			row = append(row, path+"->"+strconv.Itoa(node.Right.Val))
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
			row = append(row, path+"->"+strconv.Itoa(node.Left.Val))
		}
	}
	fmt.Println(rowList)
	return result
}
