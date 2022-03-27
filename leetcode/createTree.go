package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var bfs func(root *TreeNode) []int
	bfs = func(root *TreeNode) []int {
		nums := make([]int, 0)
		queue := list.New()
		queue.PushBack(root)
		for queue.Len() > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			nums = append(nums, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		return nums
	}

	root := createTree_iterator([]int{7, 2, 5, 6, 3, 1, 4})
	fmt.Println(bfs(root))
}

func createTree_iterator(nums []int) *TreeNode {
	nodes := make([]*TreeNode, len(nums))
	for i := 0; i < len(nums); i++ {
		nodes[i] = &TreeNode{Val: nums[i]}
	}
	for i := 0; i < len(nodes); i++ {
		if 2*i+1 < len(nodes) {
			nodes[i].Left = nodes[2*i+1]
		}
		if 2*i+2 < len(nodes) {
			nodes[i].Right = nodes[2*i+2]
		}
	}
	return nodes[0]
}

func createTree_dfs(nums []int) {

}
