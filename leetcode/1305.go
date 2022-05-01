package main

//1305. 两棵二叉搜索树中的所有元素

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	//中序遍历得到2个有序数组
	var middleOrder func(root *TreeNode, result *[]int)
	middleOrder = func(root *TreeNode, result *[]int) {
		if root == nil {
			return
		}
		middleOrder(root.Left, result)
		*result = append(*result, root.Val)
		middleOrder(root.Right, result)
	}
	arr1, arr2 := make([]int, 0), make([]int, 0)
	middleOrder(root1, &arr1)
	middleOrder(root2, &arr2)
	n1, n2 := len(arr1), len(arr2)
	result := make([]int, n1+n2)
	index := 0
	i, j := 0, 0
	for i < n1 || j < n2 {
		if i == n1 {
			result[index] = arr2[j]
			j++
		} else if j == n2 {
			result[index] = arr1[i]
			i++
		} else if arr1[i] < arr2[j] {
			result[index] = arr1[i]
			i++
		} else {
			result[index] = arr2[j]
			j++
		}
		index++
	}
	return result
}
