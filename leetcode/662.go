package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BFS 记录编号，然后最右-最左+1
func widthOfBinaryTree(root *TreeNode) int {
	max := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	//每个节点的编号，根是0，叶子是2*n+1和2*n-2
	numList := make(map[*TreeNode]int)
	numList[root] = 0
	for len(queue) > 0 {
		n := len(queue)
		width := numList[queue[n-1]] - numList[queue[0]] + 1
		if max < width {
			max = width
		}
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				numList[queue[i].Left] = numList[queue[i]]*2 + 1
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				numList[queue[i].Right] = numList[queue[i]]*2 + 2
			}
		}
		queue = queue[n:]
	}
	return max
}
